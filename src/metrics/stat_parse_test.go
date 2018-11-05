package metrics

import (
	"testing"
)

func Test_parseStats_Full(t *testing.T) {
	input := `{
		"timestamp": "2018-11-05T17:22:34",
		"MGT.uptime": {
		  "description": "Management process uptime",
		  "flag": "c", "format": "d",
		  "value": 253792
		},
		"MGT.child_start": {
		  "description": "Child process started",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MGT.child_exit": {
		  "description": "Child process normal exit",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MGT.child_stop": {
		  "description": "Child process unexpected exit",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MGT.child_died": {
		  "description": "Child process died (signal)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MGT.child_dump": {
		  "description": "Child process core dumped",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MGT.child_panic": {
		  "description": "Child process panic",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.summs": {
		  "description": "stat summ operations",
		  "flag": "c", "format": "i",
		  "value": 4
		},
		"MAIN.uptime": {
		  "description": "Child process uptime",
		  "flag": "c", "format": "d",
		  "value": 253793
		},
		"MAIN.sess_conn": {
		  "description": "Sessions accepted",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.sess_drop": {
		  "description": "Sessions dropped",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sess_fail": {
		  "description": "Session accept failures",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.client_req_400": {
		  "description": "Client requests received, subject to 400 errors",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.client_req_417": {
		  "description": "Client requests received, subject to 417 errors",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.client_req": {
		  "description": "Good client requests received",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.cache_hit": {
		  "description": "Cache hits",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.cache_hit_grace": {
		  "description": "Cache grace hits",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.cache_hitpass": {
		  "description": "Cache hits for pass.",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.cache_hitmiss": {
		  "description": "Cache hits for miss.",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.cache_miss": {
		  "description": "Cache misses",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.backend_conn": {
		  "description": "Backend conn. success",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.backend_unhealthy": {
		  "description": "Backend conn. not attempted",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.backend_busy": {
		  "description": "Backend conn. too many",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.backend_fail": {
		  "description": "Backend conn. failures",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.backend_reuse": {
		  "description": "Backend conn. reuses",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.backend_recycle": {
		  "description": "Backend conn. recycles",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.backend_retry": {
		  "description": "Backend conn. retry",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_head": {
		  "description": "Fetch no body (HEAD)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_length": {
		  "description": "Fetch with Length",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.fetch_chunked": {
		  "description": "Fetch chunked",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_eof": {
		  "description": "Fetch EOF",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_bad": {
		  "description": "Fetch bad T-E",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_none": {
		  "description": "Fetch no body",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_1xx": {
		  "description": "Fetch no body (1xx)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_204": {
		  "description": "Fetch no body (204)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_304": {
		  "description": "Fetch no body (304)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_failed": {
		  "description": "Fetch failed (all causes)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.fetch_no_thread": {
		  "description": "Fetch failed (no thread)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.pools": {
		  "description": "Number of thread pools",
		  "flag": "g", "format": "i",
		  "value": 2
		},
		"MAIN.threads": {
		  "description": "Total number of threads",
		  "flag": "g", "format": "i",
		  "value": 200
		},
		"MAIN.threads_limited": {
		  "description": "Threads hit max",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.threads_created": {
		  "description": "Threads created",
		  "flag": "c", "format": "i",
		  "value": 200
		},
		"MAIN.threads_destroyed": {
		  "description": "Threads destroyed",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.threads_failed": {
		  "description": "Thread creation failed",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.thread_queue_len": {
		  "description": "Length of session queue",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.busy_sleep": {
		  "description": "Number of requests sent to sleep on busy objhdr",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.busy_wakeup": {
		  "description": "Number of requests woken after sleep on busy objhdr",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.busy_killed": {
		  "description": "Number of requests killed after sleep on busy objhdr",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sess_queued": {
		  "description": "Sessions queued for thread",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sess_dropped": {
		  "description": "Sessions dropped for thread",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.req_dropped": {
		  "description": "Requests dropped",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.n_object": {
		  "description": "object structs made",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_vampireobject": {
		  "description": "unresurrected objects",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_objectcore": {
		  "description": "objectcore structs made",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_objecthead": {
		  "description": "objecthead structs made",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_backend": {
		  "description": "Number of backends",
		  "flag": "g", "format": "i",
		  "value": 2
		},
		"MAIN.n_expired": {
		  "description": "Number of expired objects",
		  "flag": "g", "format": "i",
		  "value": 1
		},
		"MAIN.n_lru_nuked": {
		  "description": "Number of LRU nuked objects",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_lru_moved": {
		  "description": "Number of LRU moved objects",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_lru_limited": {
		  "description": "Reached nuke_limit",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.losthdr": {
		  "description": "HTTP header overflows",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.s_sess": {
		  "description": "Total sessions seen",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.s_pipe": {
		  "description": "Total pipe sessions seen",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.s_pass": {
		  "description": "Total pass-ed requests seen",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.s_fetch": {
		  "description": "Total backend fetches initiated",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.s_synth": {
		  "description": "Total synthethic responses made",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.s_req_hdrbytes": {
		  "description": "Request header bytes",
		  "flag": "c", "format": "B",
		  "value": 373
		},
		"MAIN.s_req_bodybytes": {
		  "description": "Request body bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"MAIN.s_resp_hdrbytes": {
		  "description": "Response header bytes",
		  "flag": "c", "format": "B",
		  "value": 289
		},
		"MAIN.s_resp_bodybytes": {
		  "description": "Response body bytes",
		  "flag": "c", "format": "B",
		  "value": 2326
		},
		"MAIN.s_pipe_hdrbytes": {
		  "description": "Pipe request header bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"MAIN.s_pipe_in": {
		  "description": "Piped bytes from client",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"MAIN.s_pipe_out": {
		  "description": "Piped bytes to client",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"MAIN.sess_closed": {
		  "description": "Session Closed",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sess_closed_err": {
		  "description": "Session Closed with error",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.sess_readahead": {
		  "description": "Session Read Ahead",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sess_herd": {
		  "description": "Session herd",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.sc_rem_close": {
		  "description": "Session OK  REM_CLOSE",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_req_close": {
		  "description": "Session OK  REQ_CLOSE",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_req_http10": {
		  "description": "Session Err REQ_HTTP10",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_rx_bad": {
		  "description": "Session Err RX_BAD",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_rx_body": {
		  "description": "Session Err RX_BODY",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_rx_junk": {
		  "description": "Session Err RX_JUNK",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_rx_overflow": {
		  "description": "Session Err RX_OVERFLOW",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_rx_timeout": {
		  "description": "Session Err RX_TIMEOUT",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.sc_tx_pipe": {
		  "description": "Session OK  TX_PIPE",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_tx_error": {
		  "description": "Session Err TX_ERROR",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_tx_eof": {
		  "description": "Session OK  TX_EOF",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_resp_close": {
		  "description": "Session OK  RESP_CLOSE",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_overload": {
		  "description": "Session Err OVERLOAD",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_pipe_overflow": {
		  "description": "Session Err PIPE_OVERFLOW",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_range_short": {
		  "description": "Session Err RANGE_SHORT",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_req_http20": {
		  "description": "Session Err REQ_HTTP20",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.sc_vcl_failure": {
		  "description": "Session Err VCL_FAILURE",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.shm_records": {
		  "description": "SHM records",
		  "flag": "c", "format": "i",
		  "value": 169275
		},
		"MAIN.shm_writes": {
		  "description": "SHM writes",
		  "flag": "c", "format": "i",
		  "value": 169179
		},
		"MAIN.shm_flushes": {
		  "description": "SHM flushes due to overflow",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.shm_cont": {
		  "description": "SHM MTX contention",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.shm_cycles": {
		  "description": "SHM cycles through buffer",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.backend_req": {
		  "description": "Backend requests made",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.n_vcl": {
		  "description": "Number of loaded VCLs in total",
		  "flag": "g", "format": "i",
		  "value": 1
		},
		"MAIN.n_vcl_avail": {
		  "description": "Number of VCLs available",
		  "flag": "g", "format": "i",
		  "value": 1
		},
		"MAIN.n_vcl_discard": {
		  "description": "Number of discarded VCLs",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.vcl_fail": {
		  "description": "VCL failures",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans": {
		  "description": "Count of bans",
		  "flag": "g", "format": "i",
		  "value": 1
		},
		"MAIN.bans_completed": {
		  "description": "Number of bans marked 'completed'",
		  "flag": "g", "format": "i",
		  "value": 1
		},
		"MAIN.bans_obj": {
		  "description": "Number of bans using obj.*",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.bans_req": {
		  "description": "Number of bans using req.*",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.bans_added": {
		  "description": "Bans added",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.bans_deleted": {
		  "description": "Bans deleted",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_tested": {
		  "description": "Bans tested against objects (lookup)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_obj_killed": {
		  "description": "Objects killed by bans (lookup)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_lurker_tested": {
		  "description": "Bans tested against objects (lurker)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_tests_tested": {
		  "description": "Ban tests tested against objects (lookup)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_lurker_tests_tested": {
		  "description": "Ban tests tested against objects (lurker)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_lurker_obj_killed": {
		  "description": "Objects killed by bans (lurker)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_lurker_obj_killed_cutoff": {
		  "description": "Objects killed by bans for cutoff (lurker)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_dups": {
		  "description": "Bans superseded by other bans",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_lurker_contention": {
		  "description": "Lurker gave way for lookup",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.bans_persisted_bytes": {
		  "description": "Bytes used by the persisted ban lists",
		  "flag": "g", "format": "B",
		  "value": 16
		},
		"MAIN.bans_persisted_fragmentation": {
		  "description": "Extra bytes in persisted ban lists due to fragmentation",
		  "flag": "g", "format": "B",
		  "value": 0
		},
		"MAIN.n_purges": {
		  "description": "Number of purge operations executed",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.n_obj_purged": {
		  "description": "Number of purged objects",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MAIN.exp_mailed": {
		  "description": "Number of objects mailed to expiry thread",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.exp_received": {
		  "description": "Number of objects received by expiry thread",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.hcb_nolock": {
		  "description": "HCB Lookups without lock",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.hcb_lock": {
		  "description": "HCB Lookups with lock",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.hcb_insert": {
		  "description": "HCB Inserts",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MAIN.esi_errors": {
		  "description": "ESI parse errors (unlock)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.esi_warnings": {
		  "description": "ESI parse warnings (unlock)",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.vmods": {
		  "description": "Loaded VMODs",
		  "flag": "g", "format": "i",
		  "value": 2
		},
		"MAIN.n_gzip": {
		  "description": "Gzip operations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.n_gunzip": {
		  "description": "Gunzip operations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MAIN.n_test_gunzip": {
		  "description": "Test gunzip operations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.backend.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 3
		},
		"LCK.backend.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.backend.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 3
		},
		"LCK.ban.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.ban.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.ban.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 10238
		},
		"LCK.busyobj.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.busyobj.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.busyobj.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 10
		},
		"LCK.cli.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.cli.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.cli.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 84594
		},
		"LCK.exp.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.exp.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.exp.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 53251
		},
		"LCK.hcb.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.hcb.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.hcb.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 1412
		},
		"LCK.lru.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 2
		},
		"LCK.lru.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.lru.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 2
		},
		"LCK.mempool.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 5
		},
		"LCK.mempool.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.mempool.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 1124519
		},
		"LCK.objhdr.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 2
		},
		"LCK.objhdr.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.objhdr.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 16
		},
		"LCK.pipestat.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.pipestat.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.pipestat.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.sess.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.sess.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.sess.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 3
		},
		"LCK.tcp_pool.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 3
		},
		"LCK.tcp_pool.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.tcp_pool.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 7
		},
		"LCK.vbe.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.vbe.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.vbe.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 84587
		},
		"LCK.vcapace.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.vcapace.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.vcapace.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.vcl.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.vcl.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.vcl.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 8
		},
		"LCK.vxid.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.vxid.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.vxid.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.waiter.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 2
		},
		"LCK.waiter.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.waiter.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 5080
		},
		"LCK.wq.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 3
		},
		"LCK.wq.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.wq.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 255858
		},
		"LCK.wstat.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"LCK.wstat.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.wstat.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 59772
		},
		"MEMPOOL.busyobj.live": {
		  "description": "In use",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MEMPOOL.busyobj.pool": {
		  "description": "In Pool",
		  "flag": "g", "format": "i",
		  "value": 10
		},
		"MEMPOOL.busyobj.sz_wanted": {
		  "description": "Size requested",
		  "flag": "g", "format": "B",
		  "value": 65536
		},
		"MEMPOOL.busyobj.sz_actual": {
		  "description": "Size allocated",
		  "flag": "g", "format": "B",
		  "value": 65504
		},
		"MEMPOOL.busyobj.allocs": {
		  "description": "Allocations",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.busyobj.frees": {
		  "description": "Frees",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.busyobj.recycle": {
		  "description": "Recycled from pool",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.busyobj.timeout": {
		  "description": "Timed out from pool",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.busyobj.toosmall": {
		  "description": "Too small to recycle",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.busyobj.surplus": {
		  "description": "Too many for pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.busyobj.randry": {
		  "description": "Pool ran dry",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req0.live": {
		  "description": "In use",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req0.pool": {
		  "description": "In Pool",
		  "flag": "g", "format": "i",
		  "value": 10
		},
		"MEMPOOL.req0.sz_wanted": {
		  "description": "Size requested",
		  "flag": "g", "format": "B",
		  "value": 65536
		},
		"MEMPOOL.req0.sz_actual": {
		  "description": "Size allocated",
		  "flag": "g", "format": "B",
		  "value": 65504
		},
		"MEMPOOL.req0.allocs": {
		  "description": "Allocations",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.req0.frees": {
		  "description": "Frees",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.req0.recycle": {
		  "description": "Recycled from pool",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.req0.timeout": {
		  "description": "Timed out from pool",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.req0.toosmall": {
		  "description": "Too small to recycle",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req0.surplus": {
		  "description": "Too many for pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req0.randry": {
		  "description": "Pool ran dry",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess0.live": {
		  "description": "In use",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess0.pool": {
		  "description": "In Pool",
		  "flag": "g", "format": "i",
		  "value": 10
		},
		"MEMPOOL.sess0.sz_wanted": {
		  "description": "Size requested",
		  "flag": "g", "format": "B",
		  "value": 512
		},
		"MEMPOOL.sess0.sz_actual": {
		  "description": "Size allocated",
		  "flag": "g", "format": "B",
		  "value": 480
		},
		"MEMPOOL.sess0.allocs": {
		  "description": "Allocations",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.sess0.frees": {
		  "description": "Frees",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.sess0.recycle": {
		  "description": "Recycled from pool",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.sess0.timeout": {
		  "description": "Timed out from pool",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"MEMPOOL.sess0.toosmall": {
		  "description": "Too small to recycle",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess0.surplus": {
		  "description": "Too many for pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess0.randry": {
		  "description": "Pool ran dry",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.sma.creat": {
		  "description": "Created locks",
		  "flag": "c", "format": "i",
		  "value": 2
		},
		"LCK.sma.destroy": {
		  "description": "Destroyed locks",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"LCK.sma.locks": {
		  "description": "Lock Operations",
		  "flag": "c", "format": "i",
		  "value": 4
		},
		"SMA.s0.c_req": {
		  "description": "Allocator requests",
		  "flag": "c", "format": "i",
		  "value": 2
		},
		"SMA.s0.c_fail": {
		  "description": "Allocator failures",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"SMA.s0.c_bytes": {
		  "description": "Bytes allocated",
		  "flag": "c", "format": "B",
		  "value": 2654
		},
		"SMA.s0.c_freed": {
		  "description": "Bytes freed",
		  "flag": "c", "format": "B",
		  "value": 2654
		},
		"SMA.s0.g_alloc": {
		  "description": "Allocations outstanding",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"SMA.s0.g_bytes": {
		  "description": "Bytes outstanding",
		  "flag": "g", "format": "B",
		  "value": 0
		},
		"SMA.s0.g_space": {
		  "description": "Bytes available",
		  "flag": "g", "format": "B",
		  "value": 268435456
		},
		"SMA.Transient.c_req": {
		  "description": "Allocator requests",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"SMA.Transient.c_fail": {
		  "description": "Allocator failures",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"SMA.Transient.c_bytes": {
		  "description": "Bytes allocated",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"SMA.Transient.c_freed": {
		  "description": "Bytes freed",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"SMA.Transient.g_alloc": {
		  "description": "Allocations outstanding",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"SMA.Transient.g_bytes": {
		  "description": "Bytes outstanding",
		  "flag": "g", "format": "B",
		  "value": 0
		},
		"SMA.Transient.g_space": {
		  "description": "Bytes available",
		  "flag": "g", "format": "B",
		  "value": 0
		},
		"MEMPOOL.req1.live": {
		  "description": "In use",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.pool": {
		  "description": "In Pool",
		  "flag": "g", "format": "i",
		  "value": 10
		},
		"MEMPOOL.req1.sz_wanted": {
		  "description": "Size requested",
		  "flag": "g", "format": "B",
		  "value": 65536
		},
		"MEMPOOL.req1.sz_actual": {
		  "description": "Size allocated",
		  "flag": "g", "format": "B",
		  "value": 65504
		},
		"MEMPOOL.req1.allocs": {
		  "description": "Allocations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.frees": {
		  "description": "Frees",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.recycle": {
		  "description": "Recycled from pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.timeout": {
		  "description": "Timed out from pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.toosmall": {
		  "description": "Too small to recycle",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.surplus": {
		  "description": "Too many for pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.req1.randry": {
		  "description": "Pool ran dry",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.live": {
		  "description": "In use",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.pool": {
		  "description": "In Pool",
		  "flag": "g", "format": "i",
		  "value": 10
		},
		"MEMPOOL.sess1.sz_wanted": {
		  "description": "Size requested",
		  "flag": "g", "format": "B",
		  "value": 512
		},
		"MEMPOOL.sess1.sz_actual": {
		  "description": "Size allocated",
		  "flag": "g", "format": "B",
		  "value": 480
		},
		"MEMPOOL.sess1.allocs": {
		  "description": "Allocations",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.frees": {
		  "description": "Frees",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.recycle": {
		  "description": "Recycled from pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.timeout": {
		  "description": "Timed out from pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.toosmall": {
		  "description": "Too small to recycle",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.surplus": {
		  "description": "Too many for pool",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"MEMPOOL.sess1.randry": {
		  "description": "Pool ran dry",
		  "flag": "c", "format": "i",
		  "value": 0
		},
		"VBE.boot.web1.happy": {
		  "description": "Happy health probes",
		  "flag": "b", "format": "b",
		  "value": 0
		},
		"VBE.boot.web1.bereq_hdrbytes": {
		  "description": "Request header bytes",
		  "flag": "c", "format": "B",
		  "value": 381
		},
		"VBE.boot.web1.bereq_bodybytes": {
		  "description": "Request body bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web1.beresp_hdrbytes": {
		  "description": "Response header bytes",
		  "flag": "c", "format": "B",
		  "value": 229
		},
		"VBE.boot.web1.beresp_bodybytes": {
		  "description": "Response body bytes",
		  "flag": "c", "format": "B",
		  "value": 2326
		},
		"VBE.boot.web1.pipe_hdrbytes": {
		  "description": "Pipe request header bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web1.pipe_out": {
		  "description": "Piped bytes to backend",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web1.pipe_in": {
		  "description": "Piped bytes from backend",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web1.conn": {
		  "description": "Concurrent connections to backend",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"VBE.boot.web1.req": {
		  "description": "Backend requests sent",
		  "flag": "c", "format": "i",
		  "value": 1
		},
		"VBE.boot.web2.happy": {
		  "description": "Happy health probes",
		  "flag": "b", "format": "b",
		  "value": 0
		},
		"VBE.boot.web2.bereq_hdrbytes": {
		  "description": "Request header bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.bereq_bodybytes": {
		  "description": "Request body bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.beresp_hdrbytes": {
		  "description": "Response header bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.beresp_bodybytes": {
		  "description": "Response body bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.pipe_hdrbytes": {
		  "description": "Pipe request header bytes",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.pipe_out": {
		  "description": "Piped bytes to backend",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.pipe_in": {
		  "description": "Piped bytes from backend",
		  "flag": "c", "format": "B",
		  "value": 0
		},
		"VBE.boot.web2.conn": {
		  "description": "Concurrent connections to backend",
		  "flag": "g", "format": "i",
		  "value": 0
		},
		"VBE.boot.web2.req": {
		  "description": "Backend requests sent",
		  "flag": "c", "format": "i",
		  "value": 0
		}
	  }`

	_, _, err := parseStats([]byte(input))
	if err != nil {
		t.Fatalf("Unexpected error %s", err.Error())
	}
}
