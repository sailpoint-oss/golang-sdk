import { createRequire as __createRequire } from 'node:module';
const require = __createRequire(import.meta.url);

// src/utils/constants.ts
var OTEL_URL = "https://otel.cloud.redocly.com";
var OTEL_TRACES_URL = process.env.OTEL_TRACES_URL || `${OTEL_URL}/v1/traces`;
var DEFAULT_FETCH_TIMEOUT = 6e3;
var ANONYMOUS_ID_CACHE_FILE = "redocly-cli-anonymous-id";

export {
  OTEL_URL,
  OTEL_TRACES_URL,
  DEFAULT_FETCH_TIMEOUT,
  ANONYMOUS_ID_CACHE_FILE
};
