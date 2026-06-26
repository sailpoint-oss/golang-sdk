import { createRequire as __createRequire } from 'node:module';
const require = __createRequire(import.meta.url);
import {
  __commonJS,
  __require,
  __toESM
} from "./5ILQMFXK.js";

// ../../node_modules/@redocly/cli-otel/node_modules/ulid/dist/index.umd.js
var require_index_umd = __commonJS({
  "../../node_modules/@redocly/cli-otel/node_modules/ulid/dist/index.umd.js"(exports, module) {
    (function(global2, factory) {
      typeof exports === "object" && typeof module !== "undefined" ? factory(exports) : typeof define === "function" && define.amd ? define(["exports"], factory) : (global2 = typeof globalThis !== "undefined" ? globalThis : global2 || self, factory(global2.ULID = {}));
    })(exports, (function(exports2) {
      "use strict";
      function createError(message) {
        const err = new Error(message);
        err.source = "ulid";
        return err;
      }
      const ENCODING2 = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
      const ENCODING_LEN2 = ENCODING2.length;
      const TIME_MAX2 = Math.pow(2, 48) - 1;
      const TIME_LEN2 = 10;
      const RANDOM_LEN2 = 16;
      function replaceCharAt(str, index, char) {
        if (index > str.length - 1) {
          return str;
        }
        return str.substr(0, index) + char + str.substr(index + 1);
      }
      function incrementBase32(str) {
        let done = void 0;
        let index = str.length;
        let char;
        let charIndex;
        const maxCharIndex = ENCODING_LEN2 - 1;
        while (!done && index-- >= 0) {
          char = str[index];
          charIndex = ENCODING2.indexOf(char);
          if (charIndex === -1) {
            throw createError("incorrectly encoded string");
          }
          if (charIndex === maxCharIndex) {
            str = replaceCharAt(str, index, ENCODING2[0]);
            continue;
          }
          done = replaceCharAt(str, index, ENCODING2[charIndex + 1]);
        }
        if (typeof done === "string") {
          return done;
        }
        throw createError("cannot increment this string");
      }
      function randomChar2(prng) {
        let rand = Math.floor(prng() * ENCODING_LEN2);
        if (rand === ENCODING_LEN2) {
          rand = ENCODING_LEN2 - 1;
        }
        return ENCODING2.charAt(rand);
      }
      function encodeTime2(now, len) {
        if (isNaN(now)) {
          throw new Error(now + " must be a number");
        }
        if (now > TIME_MAX2) {
          throw createError("cannot encode time greater than " + TIME_MAX2);
        }
        if (now < 0) {
          throw createError("time must be positive");
        }
        if (Number.isInteger(Number(now)) === false) {
          throw createError("time must be an integer");
        }
        let mod;
        let str = "";
        for (; len > 0; len--) {
          mod = now % ENCODING_LEN2;
          str = ENCODING2.charAt(mod) + str;
          now = (now - mod) / ENCODING_LEN2;
        }
        return str;
      }
      function encodeRandom2(len, prng) {
        let str = "";
        for (; len > 0; len--) {
          str = randomChar2(prng) + str;
        }
        return str;
      }
      function decodeTime(id) {
        if (id.length !== TIME_LEN2 + RANDOM_LEN2) {
          throw createError("malformed ulid");
        }
        var time = id.substr(0, TIME_LEN2).split("").reverse().reduce((carry, char, index) => {
          const encodingIndex = ENCODING2.indexOf(char);
          if (encodingIndex === -1) {
            throw createError("invalid character found: " + char);
          }
          return carry += encodingIndex * Math.pow(ENCODING_LEN2, index);
        }, 0);
        if (time > TIME_MAX2) {
          throw createError("malformed ulid, timestamp too large");
        }
        return time;
      }
      function detectPrng(allowInsecure = false, root) {
        if (!root) {
          root = typeof window !== "undefined" ? window : null;
        }
        const browserCrypto = root && (root.crypto || root.msCrypto);
        if (browserCrypto) {
          return () => {
            const buffer = new Uint8Array(1);
            browserCrypto.getRandomValues(buffer);
            return buffer[0] / 255;
          };
        } else {
          try {
            const nodeCrypto = __require("crypto");
            return () => nodeCrypto.randomBytes(1).readUInt8() / 255;
          } catch (e) {
          }
        }
        if (allowInsecure) {
          try {
            console.error("secure crypto unusable, falling back to insecure Math.random()!");
          } catch (e) {
          }
          return () => Math.random();
        }
        throw createError("secure crypto unusable, insecure Math.random not allowed");
      }
      function factory(currPrng) {
        if (!currPrng) {
          currPrng = detectPrng();
        }
        return function ulid3(seedTime) {
          if (isNaN(seedTime)) {
            seedTime = Date.now();
          }
          return encodeTime2(seedTime, TIME_LEN2) + encodeRandom2(RANDOM_LEN2, currPrng);
        };
      }
      function monotonicFactory(currPrng) {
        if (!currPrng) {
          currPrng = detectPrng();
        }
        let lastTime = 0;
        let lastRandom;
        return function ulid3(seedTime) {
          if (isNaN(seedTime)) {
            seedTime = Date.now();
          }
          if (seedTime <= lastTime) {
            const incrementedRandom = lastRandom = incrementBase32(lastRandom);
            return encodeTime2(lastTime, TIME_LEN2) + incrementedRandom;
          }
          lastTime = seedTime;
          const newRandom = lastRandom = encodeRandom2(RANDOM_LEN2, currPrng);
          return encodeTime2(seedTime, TIME_LEN2) + newRandom;
        };
      }
      const ulid2 = factory();
      exports2.decodeTime = decodeTime;
      exports2.detectPrng = detectPrng;
      exports2.encodeRandom = encodeRandom2;
      exports2.encodeTime = encodeTime2;
      exports2.factory = factory;
      exports2.incrementBase32 = incrementBase32;
      exports2.monotonicFactory = monotonicFactory;
      exports2.randomChar = randomChar2;
      exports2.replaceCharAt = replaceCharAt;
      exports2.ulid = ulid2;
    }));
  }
});

// ../../node_modules/@redocly/cli-otel/lib/index.js
var import_ulid = __toESM(require_index_umd(), 1);
var z = Object.defineProperty;
var N = (e, n) => {
  for (var o in n) z(e, o, { get: n[o], enumerable: true });
};
var j = {};
N(j, { CLOUD_EVENT_STANDARD_KEYS: () => b, mapToCloudEvent: () => q });
var b = ["id", "specversion", "object", "datacontenttype", "type", "time", "origin", "env", "category", "signal", "source", "subject", "subjects", "data", "actor", "requestId", "clientIp", "organizationId", "organizationSlug", "projectId", "projectSlug", "osPlatform", "userAgent", "sessionId"];
function q(e) {
  let { type: n, data: o, actor: t, requestId: s, clientIp: r, origin: i, env: c, source: a, organizationId: m, organizationSlug: C, projectId: A, projectSlug: _, category: k = "product", signal: D = "log", osPlatform: h, userAgent: w, sessionId: P } = e, $ = !!(o && Array.isArray(o)), v = t ? { id: t.id ?? `ann_${(0, import_ulid.ulid)()}`, object: t.object ?? "user", uri: t.uri ?? "" } : null, E = `evt_${(0, import_ulid.ulid)()}`.toLowerCase(), x = /* @__PURE__ */ new Date(), I = a ?? v?.uri ?? null, T = { ...m !== void 0 && { organizationId: m }, ...C !== void 0 && { organizationSlug: C }, ...A !== void 0 && { projectId: A }, ..._ !== void 0 && { projectSlug: _ } }, g = { specversion: "1.0", object: "event", datacontenttype: "application/json; charset=utf-8", origin: i, env: c, category: k, signal: D, osPlatform: h, userAgent: w, clientIp: r, sessionId: P };
  if ($) {
    let u = o, B = u.map((y) => ({ id: y.id ?? "", object: y.object ?? "", uri: y.uri ?? "" })), U = u[0]?.id ?? null;
    Object.assign(g, { ...T, id: E, type: n, time: x, source: I, actor: v, subject: U, subjects: B, data: u, requestId: s ?? "" });
  } else {
    let u = (Array.isArray(o) ? o[0] : o)?.id ?? null;
    Object.assign(g, { ...T, id: E, type: n, time: x, source: I, actor: v, subject: u, requestId: s ?? "", data: o, ...a != null && { request: { source: a } } });
  }
  return g;
}
var O = "context";
var d = [O];
function f(e) {
  return e.replace(/([A-Z])/g, "_$1").toLowerCase();
}
function S(e) {
  let n = { updated: [], notUpdated: [] };
  for (let o of e) {
    if (!o || typeof o != "object") continue;
    let t = o.object;
    if (typeof t != "string") continue;
    let s = o, r = n.notUpdated.find((i) => i.object === t);
    r ? (n.updated = [r, s], n.notUpdated = n.notUpdated.filter((i) => i.object !== t)) : n.updated.some((i) => i.object === t) || n.notUpdated.push(s);
  }
  return n;
}
function p(e, n, o) {
  let t = o ? `${o}.` : "", s = n === "" ? "" : `${n}.`;
  return `${e}.${t}${s}`;
}
function R(e, n) {
  let o = e[O];
  return o === "before" || o === "after" ? o : n === 0 ? "after" : "before";
}
function l(e, n, o, t = true, s) {
  let r = s?.length ? new Set(s) : null;
  for (let [i, c] of Object.entries(e)) if (!r?.has(i) && c !== void 0) {
    let a = t ? f(i) : i;
    o[`${n}${a}`] = c;
  }
}
function W(e, n, o) {
  for (let [t, s] of Object.entries(e)) s == null || typeof s == "object" || (o[`${n}.${f(t)}`] = s);
}
function K(e, n, o) {
  let t = n instanceof Date ? n.toISOString() : new Date(n).toISOString();
  return { "cloudevents.event_id": e.id, "cloudevents.event_type": e.type, "cloudevents.event_source": e.source ?? void 0, "cloudevents.event_spec_version": e.specversion, "cloudevents.event_data_content_type": e.datacontenttype ?? "application/json; charset=utf-8", "cloudevents.event_time": t, "cloudevents.event_subject": e.subject ?? "", "cloudevents.page.uri": typeof location < "u" ? location.href : void 0, "cloudevents.event_version": o?.version, "cloudevents.event_origin": e.origin ?? o?.serviceName, "cloudevents.event_env": e.env, "cloudevents.event_source_details.id": e.actor?.id ?? "anonymous", "cloudevents.event_source_details.object": e.actor?.object ?? "anonymous", "cloudevents.event_source_details.uri": e.actor?.uri ?? void 0, "cloudevents.event_client_ip": e.clientIp, "cloudevents.event_object": e.object || "event", "cloudevents.event_category": e.category, "cloudevents.event_signal": e.signal, "cloudevents.event_actor.id": e.actor?.id ?? void 0, "cloudevents.event_actor.object": e.actor?.object ?? void 0, "cloudevents.event_actor.uri": e.actor?.uri ?? void 0, "cloudevents.event_organization_id": e.organizationId, "cloudevents.event_organization_slug": e.organizationSlug, "cloudevents.event_project_id": e.projectId, "cloudevents.event_project_slug": e.projectSlug, "cloudevents.event_request_id": e.requestId, "cloudevents.event_os_platform": e.osPlatform, "cloudevents.event_user_agent": e.userAgent, "cloudevents.event_session_id": e.sessionId };
}
function G(e, n, o) {
  let t = e.data, s = { ...K(e, n, o) };
  return L(t, s), X(e, s), Z(e, s), s;
}
function L(e, n) {
  !e || typeof e != "object" || (Array.isArray(e) ? V(e, n) : H(e, n));
}
function V(e, n) {
  e.some((t) => t && typeof t == "object" && typeof t.object == "string") ? Y(e, n) : F(e, n);
}
function Y(e, n) {
  let o = S(e);
  o.updated.forEach((t, s) => {
    let r = p("cloudevents.event_data", t.object, R(t, s));
    l(t, r, n, true, d);
  }), o.notUpdated.forEach((t) => {
    let s = p("cloudevents.event_data", t.object, void 0);
    l(t, s, n, true, d);
  });
}
function F(e, n) {
  for (let o of e) if (!(!o || typeof o != "object")) for (let [t, s] of Object.entries(o)) {
    if (s === void 0) continue;
    let r = f(t);
    s !== null && typeof s == "object" && !Array.isArray(s) ? W(s, `cloudevents.event_data.${r}`, n) : n[`cloudevents.event_data.${r}`] = s;
  }
}
function H(e, n) {
  for (let [o, t] of Object.entries(e)) {
    if (t == null) continue;
    let s = f(o);
    if (typeof t == "object" && !Array.isArray(t)) for (let [r, i] of Object.entries(t)) i != null && (typeof i == "string" || typeof i == "number" || typeof i == "boolean") && (n[`cloudevents.event_data.${s}.${r}`] = i);
    else (typeof t == "string" || typeof t == "number" || typeof t == "boolean") && (n[`cloudevents.event_data.${s}`] = t);
  }
}
function X(e, n) {
  if (!e.subjects || !Array.isArray(e.subjects)) return;
  let o = S(e.subjects);
  o.updated.forEach((t, s) => {
    let r = p("cloudevents.event_subjects", t.object, R(t, s));
    l(t, r, n, false, d);
  }), o.notUpdated.forEach((t) => {
    let s = p("cloudevents.event_subjects", t.object, void 0);
    l(t, s, n, false, d);
  });
}
function Z(e, n) {
  for (let [o, t] of Object.entries(e)) if (!(b.includes(o) || t === void 0)) if (t instanceof Object) for (let [s, r] of Object.entries(t)) r !== void 0 && (n[`cloudevents.${o}.${s}`] = r);
  else n[`cloudevents.${o}`] = t;
}

// ../../node_modules/ulid/dist/node/index.js
import crypto from "node:crypto";
var ENCODING = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";
var ENCODING_LEN = 32;
var RANDOM_LEN = 16;
var TIME_LEN = 10;
var TIME_MAX = 281474976710655;
var ULIDErrorCode;
(function(ULIDErrorCode2) {
  ULIDErrorCode2["Base32IncorrectEncoding"] = "B32_ENC_INVALID";
  ULIDErrorCode2["DecodeTimeInvalidCharacter"] = "DEC_TIME_CHAR";
  ULIDErrorCode2["DecodeTimeValueMalformed"] = "DEC_TIME_MALFORMED";
  ULIDErrorCode2["EncodeTimeNegative"] = "ENC_TIME_NEG";
  ULIDErrorCode2["EncodeTimeSizeExceeded"] = "ENC_TIME_SIZE_EXCEED";
  ULIDErrorCode2["EncodeTimeValueMalformed"] = "ENC_TIME_MALFORMED";
  ULIDErrorCode2["PRNGDetectFailure"] = "PRNG_DETECT";
  ULIDErrorCode2["ULIDInvalid"] = "ULID_INVALID";
  ULIDErrorCode2["Unexpected"] = "UNEXPECTED";
  ULIDErrorCode2["UUIDInvalid"] = "UUID_INVALID";
})(ULIDErrorCode || (ULIDErrorCode = {}));
var ULIDError = class extends Error {
  constructor(errorCode, message) {
    super(`${message} (${errorCode})`);
    this.name = "ULIDError";
    this.code = errorCode;
  }
};
function randomChar(prng) {
  const randomPosition = Math.floor(prng() * ENCODING_LEN) % ENCODING_LEN;
  return ENCODING.charAt(randomPosition);
}
function detectPRNG(root) {
  const rootLookup = detectRoot();
  const globalCrypto = rootLookup && (rootLookup.crypto || rootLookup.msCrypto) || (typeof crypto !== "undefined" ? crypto : null);
  if (typeof globalCrypto?.getRandomValues === "function") {
    return () => {
      const buffer = new Uint8Array(1);
      globalCrypto.getRandomValues(buffer);
      return buffer[0] / 256;
    };
  } else if (typeof globalCrypto?.randomBytes === "function") {
    return () => globalCrypto.randomBytes(1).readUInt8() / 256;
  } else if (crypto?.randomBytes) {
    return () => crypto.randomBytes(1).readUInt8() / 256;
  }
  throw new ULIDError(ULIDErrorCode.PRNGDetectFailure, "Failed to find a reliable PRNG");
}
function detectRoot() {
  if (inWebWorker())
    return self;
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  return null;
}
function encodeRandom(len, prng) {
  let str = "";
  for (; len > 0; len--) {
    str = randomChar(prng) + str;
  }
  return str;
}
function encodeTime(now, len = TIME_LEN) {
  if (isNaN(now)) {
    throw new ULIDError(ULIDErrorCode.EncodeTimeValueMalformed, `Time must be a number: ${now}`);
  } else if (now > TIME_MAX) {
    throw new ULIDError(ULIDErrorCode.EncodeTimeSizeExceeded, `Cannot encode a time larger than ${TIME_MAX}: ${now}`);
  } else if (now < 0) {
    throw new ULIDError(ULIDErrorCode.EncodeTimeNegative, `Time must be positive: ${now}`);
  } else if (Number.isInteger(now) === false) {
    throw new ULIDError(ULIDErrorCode.EncodeTimeValueMalformed, `Time must be an integer: ${now}`);
  }
  let mod, str = "";
  for (let currentLen = len; currentLen > 0; currentLen--) {
    mod = now % ENCODING_LEN;
    str = ENCODING.charAt(mod) + str;
    now = (now - mod) / ENCODING_LEN;
  }
  return str;
}
function inWebWorker() {
  return typeof WorkerGlobalScope !== "undefined" && self instanceof WorkerGlobalScope;
}
function ulid(seedTime, prng) {
  const currentPRNG = prng || detectPRNG();
  const seed = !seedTime || isNaN(seedTime) ? Date.now() : seedTime;
  return encodeTime(seed, TIME_LEN) + encodeRandom(RANDOM_LEN, currentPRNG);
}

export {
  j,
  G,
  ulid
};
