#include "wiredtiger.h"

#include <iostream>
#include <sstream>

class WtSession;
void invariantFailedWithLines(int ret, const char* file, unsigned line);
void invariantFailedWithLines(WT_SESSION* session, int ret, const char* file, unsigned line);
void invariantFailedWithLines(WtSession& session, int ret, const char* file, unsigned line);

#define invariantWTOKEasy(fn)                                   \
    {                                                           \
        int _ret = fn;                                          \
        if (_ret) {                                             \
            invariantFailedWithLines(_ret, __FILE__, __LINE__); \
        }                                                       \
    }

#define invariantWTOK(session, fn)                                       \
    {                                                                    \
        int _ret = fn;                                                   \
        if (_ret) {                                                      \
            invariantFailedWithLines(session, _ret, __FILE__, __LINE__); \
        }                                                                \
    }

#define invariantWTFail(session, fn)                                     \
    {                                                                    \
        int _ret = fn;                                                   \
        if (!_ret) {                                                     \
            invariantFailedWithLines(session, _ret, __FILE__, __LINE__); \
        }                                                                \
    }

class WtCursor {
public:
    WtCursor(WT_SESSION* session, std::string uri, std::string config) {
        invariantWTOK(
            session, session->open_cursor(session, uri.c_str(), nullptr, config.c_str(), &_cursor));
    }

    WtCursor(WtCursor&& other) : _cursor(other._cursor) {
        other._cursor = nullptr;
    }

    WtCursor& operator=(WtCursor&& other) {
        _cursor = other._cursor;
        other._cursor = nullptr;
        return *this;
    }

    ~WtCursor() {
        if (_cursor) {
            invariantWTOK(_cursor->session, _cursor->close(_cursor));
        }
    }

    WT_CURSOR* operator->() const {
        return _cursor;
    }

    WT_CURSOR* get() {
        return _cursor;
    }

    /**
     * @return true on more entries.
     */
    bool next() {
        return !_cursor->next(_cursor);
    }

    bool prev() {
        return !_cursor->prev(_cursor);
    }

    void setKey(int64_t key) {
        _cursor->set_key(_cursor, key);
    }

    void setValue(int64_t value) {
        _cursor->set_value(_cursor, value);
    }

    void setByteValue(char* value) {
        _cursor->set_value(_cursor, value);
    }

    void remove() {
        invariantWTOKEasy(_cursor->remove(_cursor));
    }

    int64_t getKey() {
        int64_t ret;
        invariantWTOKEasy(_cursor->get_key(_cursor, &ret));
        return ret;
    }

    std::string getKeyString() {
        char* filename;
        invariantWTOKEasy(_cursor->get_key(_cursor, &filename));
        return std::string(filename);
    }

    int64_t getIntValue() {
        int64_t ret;
        invariantWTOKEasy(_cursor->get_value(_cursor, &ret));
        return ret;
    }

    char* getByteValue(int64_t key) {
        setKey(key);
        if (_cursor->search(_cursor) == WT_NOTFOUND) {
            std::cout << "\tKey not found, returning -1. Key: " << key << std::endl;
            return nullptr;
        }

        char* ret;
        invariantWTOK(_cursor->session, _cursor->get_value(_cursor, &ret));
        return ret;
    }

    int save() {
        return _cursor->insert(_cursor);
    }

    int insert(int64_t key, int64_t value) {
        setKey(key);
        setValue(value);
        return save();
    }

    int searchExact(int64_t key) {
        int64_t ret;
        setKey(key);
        int err = _cursor->search(_cursor);
        if (err == WT_NOTFOUND) {
            // std::cout << "\tKey not found, returning -1. Key: " << key << std::endl;
            return -1;
        }
        if (err) {
            return err;
        }
        invariantWTOK(_cursor->session, _cursor->get_value(_cursor, &ret));
        return ret;
    }

    int remove(int64_t key) {
        setKey(key);
        return _cursor->remove(_cursor);
    }

private:
    WT_CURSOR* _cursor;
};

class WtSession {
public:
    WtSession(WT_CONNECTION* conn) {
        if (conn->open_session(conn, nullptr, "isolation=snapshot", &_session)) {
            throw "Session open failed";
        }
    }

    WtSession(WtSession&& other) : _session(other._session) {
        other._session = nullptr;
    }

    WtSession& operator=(WtSession&& other) {
        _session = other._session;
        other._session = nullptr;
        return *this;
    }

    ~WtSession() {
        if (_session) {
            invariantWTOK(_session, _session->close(_session, nullptr));
        }
    }

    WT_SESSION* operator->() const {
        return _session;
    }

    WT_SESSION* get() const {
        return _session;
    }

    int createTable(std::string uri, std::string conf = "") {
        std::stringstream ss;
        ss << "key_format=q,value_format=q,log=(enabled=false)," << conf;
        return _session->create(_session, uri.c_str(), ss.str().c_str());
    }

    int createByteValueTable(std::string uri) {
        return _session->create(
            _session, uri.c_str(), "key_format=q,value_format=u,log=(enabled=false)");
    }

    int createTableWithLog(std::string uri) {
        return _session->create(
            _session, uri.c_str(), "key_format=q,value_format=q,log=(enabled=true)");
    }

    void alterTableLogging(std::string uri, bool enable) {
        if (enable) {
            invariantWTOK(_session, _session->alter(_session, uri.c_str(), "log=(enabled=true)"));
        } else {
            invariantWTOK(_session, _session->alter(_session, uri.c_str(), "log=(enabled=false)"));
        }
    }

    void alterTableCommitTimestampAssertions(std::string uri, std::string verificationValue) {
        invariantWTOK(
            _session,
            _session->alter(_session,
                            uri.c_str(),
                            ("assert=(commit_timestamp=" + verificationValue + ")").c_str()));
    }

    /**
     * @param required can be [always, never, none].
     */
    int alterTableRequireCommitTimestamps(std::string uri, std::string required) {
        std::stringstream ss;
        ss << "assert=(commit_timestamp=" << required << ")";
        return _session->alter(_session, uri.c_str(), ss.str().c_str());
    }

    int dropTable(std::string uri) {
        return _session->drop(_session, uri.c_str(), nullptr);
    }

    WtCursor openCursor(std::string uri, std::string config = "") {
        return WtCursor(_session, uri, config);
    }

    WtCursor openBackupCursor() {
        return WtCursor(_session, "backup:", "");
    }

    int readAtTimestamp(std::uint64_t timestamp) {
        char readTSConfigString[15 /* read_timestamp= */ + 16 /* 16 hexadecimal digits */ +
                                17 /* ,round_to_oldest= */ + 5 /* false */ + 1 /* trailing null */];
        auto size = std::snprintf(readTSConfigString,
                                  sizeof(readTSConfigString),
                                  "read_timestamp=%llx,round_to_oldest=false",
                                  static_cast<unsigned long long>(timestamp));
        return _session->timestamp_transaction(_session, readTSConfigString);
    }

    int beginAtTimestamp(std::uint64_t timestamp) {
        char readTSConfigString[15 /* read_timestamp= */ + (8 * 2) /* 8 hexadecimal characters */ +
                                1 /* trailing null */];
        auto size = std::snprintf(readTSConfigString,
                                  sizeof(readTSConfigString),
                                  "read_timestamp=%llx",
                                  static_cast<unsigned long long>(timestamp));
        invariantWTOK(_session, _session->begin_transaction(_session, readTSConfigString));
        return 0;
    }

    int begin(bool ignorePrepare = false) {
        if (ignorePrepare) {
            return _session->begin_transaction(_session, "ignore_prepare=true");
        } else {
            return _session->begin_transaction(_session, nullptr);
        }
    }

    int rollback() {
        return _session->rollback_transaction(_session, nullptr);
    }

    int commit(int time = 0) {
        if (time) {
            std::stringstream ss;
            ss << "commit_timestamp=" << std::hex << time;
            const std::string conf = ss.str();
            return _session->commit_transaction(_session, conf.c_str());
        } else {
            return _session->commit_transaction(_session, nullptr);
        }
    }

    int prepare(int time) {
        std::stringstream ss;
        ss << "prepare_timestamp=" << std::hex << time;
        const std::string conf = ss.str();
        int ret = _session->prepare_transaction(_session, conf.c_str());
        std::cout << "PrepRet: " << ret << std::endl;
        return ret;
    }

    int setTimestamp(int time) {
        std::stringstream ss;
        ss << "commit_timestamp=" << std::hex << time;
        const std::string conf = ss.str();
        return _session->timestamp_transaction(_session, conf.c_str());
    }

    int truncateAfter(WtCursor& start) {
        return _session->truncate(_session, nullptr, start.get(), nullptr, nullptr);
    }

private:
    WT_SESSION* _session;
};

class MetadataCursor {
public:
    MetadataCursor(WtSession&& session) : _session(std::move(session)) {
        invariantWTOK(
            _session,
            _session->open_cursor(_session.get(), "metadata:", nullptr, nullptr, &_cursor));
    }

    MetadataCursor operator=(const MetadataCursor& other) = delete;

    MetadataCursor(MetadataCursor&& other)
        : _session(std::move(other._session)), _cursor(nullptr) {}

    /**
     * @return true on more entries.
     */
    bool next() {
        return !_cursor->next(_cursor);
    }

    const char* getKey() {
        const char* ret;
        invariantWTOK(_session, _cursor->get_key(_cursor, &ret));
        return ret;
    }

    const char* getValue() {
        const char* ret;
        invariantWTOK(_session, _cursor->get_value(_cursor, &ret));
        return ret;
    }

private:
    WtSession _session;
    WT_CURSOR* _cursor;
};

class WtConn {
public:
    WtConn(std::string dbpath, std::string config = "", bool stableCheckpointOnClose = true)
        : _stableCheckpointOnClose(stableCheckpointOnClose) {
        std::stringstream ss;
        ss << "create,cache_size=1GB,log=(enabled=true)";
        if (config.length() > 0) {
            ss << "," << config;
        }

        _handler.handle_error = nullptr;
        _handler.handle_message = nullptr;
        _handler.handle_progress = nullptr;
        _handler.handle_close = nullptr;
        if (wiredtiger_open(dbpath.c_str(), &_handler, ss.str().c_str(), &_conn)) {
            throw "wiredtiger_open failed";
        }
    }

    WtConn(WtConn& other) = delete;

    ~WtConn() {
        if (_stableCheckpointOnClose) {
            // std::cout << "Using timestamps on close." << std::endl;
            if (_conn->close(_conn, "use_timestamp=true")) {
                // throw "wiredtiger_close failed";
            }
        } else {
            if (_conn->close(_conn, "use_timestamp=false")) {
                // throw "wiredtiger_close failed";
            }
        }
    }

    WtSession getSession() {
        return WtSession(_conn);
    }

    int stableCheckpoint() {
        WtSession session = getSession();
        return session->checkpoint(session.get(), "use_timestamp=true");
    }

    int unstableCheckpoint() {
        WtSession session = getSession();
        return session->checkpoint(session.get(), "use_timestamp=false");
    }

    int setStableTimestamp(std::uint64_t stableTimestamp) {
        char stableTSConfigString[17 /* "stable_timestamp= */ +
                                  (8 * 2) /* 16 hexadecimal digits */ + 1 /* trailing null */];
        auto size = std::snprintf(stableTSConfigString,
                                  sizeof(stableTSConfigString),
                                  "stable_timestamp=%llx",
                                  static_cast<unsigned long long>(stableTimestamp));
        return _conn->set_timestamp(_conn, stableTSConfigString);
    }

    int setOldestTimestamp(std::uint64_t oldestTimestamp) {
        char oldestTSConfigString[17 /* "oldest_timestamp= */ +
                                  (8 * 2) /* 16 hexadecimal digits */ + 1 /* trailing null */];
        auto size = std::snprintf(oldestTSConfigString,
                                  sizeof(oldestTSConfigString),
                                  "oldest_timestamp=%llx",
                                  static_cast<unsigned long long>(oldestTimestamp));
        return _conn->set_timestamp(_conn, oldestTSConfigString);
    }

    MetadataCursor getMetadataCursor() {
        return MetadataCursor(getSession());
    }


    uint64_t queryCommittedTimestamp() {
        char buf[(2 * 8 /*bytes in hex*/) + 1 /*nul terminator*/];
        invariantWTOKEasy(_conn->query_timestamp(_conn, buf, "get=all_committed"));

        return static_cast<uint64_t>(strtol(buf, nullptr, 16));
    }

    uint64_t queryRecoveryTimestamp() {
        char buf[(2 * 8 /*bytes in hex*/) + 1 /*nul terminator*/];
        invariantWTOKEasy(_conn->query_timestamp(_conn, buf, "get=recovery"));

        return static_cast<uint64_t>(strtol(buf, nullptr, 16));
    }

    uint64_t queryCheckpointTimestamp() {
        char buf[(2 * 8 /*bytes in hex*/) + 1 /*nul terminator*/];
        invariantWTOKEasy(_conn->query_timestamp(_conn, buf, "get=last_checkpoint"));

        return static_cast<uint64_t>(strtol(buf, nullptr, 16));
    }

    int rollbackToStable() {
        return _conn->rollback_to_stable(_conn, nullptr);
    }

    int setRelease(std::string releaseValue) {
        std::string configStr = "compatibility=(release=" + releaseValue + ")";
        return _conn->reconfigure(_conn, configStr.c_str());
    }

    /**
     * @param toDebug is any of "sessions,cursors,handles,log,txn". E.g:
     *                "sessions=true,cursors=true".
     */
    void debugInfo(std::string toDebug) {
        _conn->debug_info(_conn, toDebug.c_str());
    }

    void readHelper(std::string tableUri, int key, int timestamp = 0) {
        WtSession session = getSession();
        if (timestamp > 0) {
            session.readAtTimestamp(timestamp);
        }

        WtCursor cursor = session.openCursor(tableUri);
        std::cout << "Reading. Table: " << tableUri << " Key: " << key
                  << " Timestamp: " << timestamp << " Value: " << cursor.searchExact(key)
                  << std::endl;
    }

private:
    WT_CONNECTION* _conn;
    const bool _stableCheckpointOnClose;
    WT_EVENT_HANDLER _handler;
};

void invariantFailedWithLines(int ret, const char* file, unsigned line) {
    std::cout << file << ':' << line << ": Error: " << ret
              << " Message: " << wiredtiger_strerror(ret) << std::endl;
    exit(0);
}

void invariantFailedWithLines(WT_SESSION* session, int ret, const char* file, unsigned line) {
    std::cout << file << ':' << line << ": Error: " << ret
              << " Message: " << session->strerror(session, ret) << std::endl;
    exit(0);
}

void invariantFailedWithLines(WtSession& session, int ret, const char* file, unsigned line) {
    std::cout << file << ':' << line << ": Error: " << ret
              << " Message: " << session->strerror(session.get(), ret) << std::endl;
    exit(0);
}
