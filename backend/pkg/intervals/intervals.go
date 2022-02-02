package intervals

const EVENTS_COMMIT_INTERVAL = 30 * 1000
const HEARTBEAT_INTERVAL = 2 * 60 * 1000
const INTEGRATIONS_REQUEST_INTERVAL = 1 * 60 * 1000
const EVENTS_PAGE_EVENT_TIMEOUT = 2 * 60 * 1000
const EVENTS_INPUT_EVENT_TIMEOUT = 2 * 60 * 1000
const EVENTS_PERFORMANCE_AGGREGATION_TIMEOUT =  2 * 60 * 1000
const EVENTS_SESSION_END_TIMEOUT = HEARTBEAT_INTERVAL + 30 * 1000
const EVENTS_SESSION_END_TIMEOUT_WITH_INTEGRATIONS = HEARTBEAT_INTERVAL + 3 * 60 * 1000
const EVENTS_BACK_COMMIT_GAP = EVENTS_SESSION_END_TIMEOUT_WITH_INTEGRATIONS + 1*60*1000
