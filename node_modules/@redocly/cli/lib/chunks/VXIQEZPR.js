import { createRequire as __createRequire } from 'node:module';
const require = __createRequire(import.meta.url);
import {
  __commonJS,
  __require
} from "./5ILQMFXK.js";

// ../../node_modules/foreach/index.js
var require_foreach = __commonJS({
  "../../node_modules/foreach/index.js"(exports, module) {
    var hasOwn = Object.prototype.hasOwnProperty;
    var toString = Object.prototype.toString;
    module.exports = function forEach(obj, fn, ctx) {
      if (toString.call(fn) !== "[object Function]") {
        throw new TypeError("iterator must be a function");
      }
      var l = obj.length;
      if (l === +l) {
        for (var i = 0; i < l; i++) {
          fn.call(ctx, obj[i], i, obj);
        }
      } else {
        for (var k in obj) {
          if (hasOwn.call(obj, k)) {
            fn.call(ctx, obj[k], k, obj);
          }
        }
      }
    };
  }
});

// ../../node_modules/json-pointer/index.js
var require_json_pointer = __commonJS({
  "../../node_modules/json-pointer/index.js"(exports, module) {
    "use strict";
    var each = require_foreach();
    module.exports = api;
    function api(obj, pointer, value) {
      if (arguments.length === 3) {
        return api.set(obj, pointer, value);
      }
      if (arguments.length === 2) {
        return api.get(obj, pointer);
      }
      var wrapped = api.bind(api, obj);
      for (var name in api) {
        if (api.hasOwnProperty(name)) {
          wrapped[name] = api[name].bind(wrapped, obj);
        }
      }
      return wrapped;
    }
    api.get = function get(obj, pointer) {
      var refTokens = Array.isArray(pointer) ? pointer : api.parse(pointer);
      for (var i = 0; i < refTokens.length; ++i) {
        var tok = refTokens[i];
        if (!(typeof obj == "object" && tok in obj)) {
          throw new Error("Invalid reference token: " + tok);
        }
        obj = obj[tok];
      }
      return obj;
    };
    api.set = function set(obj, pointer, value) {
      var refTokens = Array.isArray(pointer) ? pointer : api.parse(pointer), nextTok = refTokens[0];
      if (refTokens.length === 0) {
        throw Error("Can not set the root object");
      }
      for (var i = 0; i < refTokens.length - 1; ++i) {
        var tok = refTokens[i];
        if (typeof tok !== "string" && typeof tok !== "number") {
          tok = String(tok);
        }
        if (tok === "__proto__" || tok === "constructor" || tok === "prototype") {
          continue;
        }
        if (tok === "-" && Array.isArray(obj)) {
          tok = obj.length;
        }
        nextTok = refTokens[i + 1];
        if (!(tok in obj)) {
          if (nextTok.match(/^(\d+|-)$/)) {
            obj[tok] = [];
          } else {
            obj[tok] = {};
          }
        }
        obj = obj[tok];
      }
      if (nextTok === "-" && Array.isArray(obj)) {
        nextTok = obj.length;
      }
      obj[nextTok] = value;
      return this;
    };
    api.remove = function(obj, pointer) {
      var refTokens = Array.isArray(pointer) ? pointer : api.parse(pointer);
      var finalToken = refTokens[refTokens.length - 1];
      if (finalToken === void 0) {
        throw new Error('Invalid JSON pointer for remove: "' + pointer + '"');
      }
      var parent = api.get(obj, refTokens.slice(0, -1));
      if (Array.isArray(parent)) {
        var index = +finalToken;
        if (finalToken === "" && isNaN(index)) {
          throw new Error('Invalid array index: "' + finalToken + '"');
        }
        Array.prototype.splice.call(parent, index, 1);
      } else {
        delete parent[finalToken];
      }
    };
    api.dict = function dict(obj, descend) {
      var results = {};
      api.walk(obj, function(value, pointer) {
        results[pointer] = value;
      }, descend);
      return results;
    };
    api.walk = function walk(obj, iterator, descend) {
      var refTokens = [];
      descend = descend || function(value) {
        var type = Object.prototype.toString.call(value);
        return type === "[object Object]" || type === "[object Array]";
      };
      (function next(cur) {
        each(cur, function(value, key) {
          refTokens.push(String(key));
          if (descend(value)) {
            next(value);
          } else {
            iterator(value, api.compile(refTokens));
          }
          refTokens.pop();
        });
      })(obj);
    };
    api.has = function has(obj, pointer) {
      try {
        api.get(obj, pointer);
      } catch (e) {
        return false;
      }
      return true;
    };
    api.escape = function escape(str) {
      return str.toString().replace(/~/g, "~0").replace(/\//g, "~1");
    };
    api.unescape = function unescape(str) {
      return str.replace(/~1/g, "/").replace(/~0/g, "~");
    };
    api.parse = function parse(pointer) {
      if (pointer === "") {
        return [];
      }
      if (pointer.charAt(0) !== "/") {
        throw new Error("Invalid JSON pointer: " + pointer);
      }
      return pointer.substring(1).split(/\//).map(api.unescape);
    };
    api.compile = function compile(refTokens) {
      if (refTokens.length === 0) {
        return "";
      }
      return "/" + refTokens.map(api.escape).join("/");
    };
  }
});

// ../../node_modules/openapi-sampler/dist/openapi-sampler.js
var require_openapi_sampler = __commonJS({
  "../../node_modules/openapi-sampler/dist/openapi-sampler.js"(exports, module) {
    (function(f) {
      if (typeof exports === "object" && typeof module !== "undefined") {
        module.exports = f();
      } else if (typeof define === "function" && define.amd) {
        define([], f);
      } else {
        var g;
        if (typeof window !== "undefined") {
          g = window;
        } else if (typeof global !== "undefined") {
          g = global;
        } else if (typeof self !== "undefined") {
          g = self;
        } else {
          g = this;
        }
        g.OpenAPISampler = f();
      }
    })(function() {
      var define2, module2, exports2;
      return (/* @__PURE__ */ (function() {
        function r(e, n, t) {
          function o(i2, f) {
            if (!n[i2]) {
              if (!e[i2]) {
                var c = "function" == typeof __require && __require;
                if (!f && c) return c(i2, true);
                if (u) return u(i2, true);
                var a = new Error("Cannot find module '" + i2 + "'");
                throw a.code = "MODULE_NOT_FOUND", a;
              }
              var p = n[i2] = { exports: {} };
              e[i2][0].call(p.exports, function(r2) {
                var n2 = e[i2][1][r2];
                return o(n2 || r2);
              }, p, p.exports, r, e, n, t);
            }
            return n[i2].exports;
          }
          for (var u = "function" == typeof __require && __require, i = 0; i < t.length; i++) o(t[i]);
          return o;
        }
        return r;
      })())({ 1: [function(require2, module3, exports3) {
        (() => {
          "use strict";
          var t = { d: (e2, n2) => {
            for (var i2 in n2) t.o(n2, i2) && !t.o(e2, i2) && Object.defineProperty(e2, i2, { enumerable: true, get: n2[i2] });
          }, o: (t2, e2) => Object.prototype.hasOwnProperty.call(t2, e2), r: (t2) => {
            "undefined" != typeof Symbol && Symbol.toStringTag && Object.defineProperty(t2, Symbol.toStringTag, { value: "Module" }), Object.defineProperty(t2, "__esModule", { value: true });
          } }, e = {};
          t.r(e), t.d(e, { XMLBuilder: () => gt, XMLParser: () => it, XMLValidator: () => xt });
          const n = ":A-Za-z_\\u00C0-\\u00D6\\u00D8-\\u00F6\\u00F8-\\u02FF\\u0370-\\u037D\\u037F-\\u1FFF\\u200C-\\u200D\\u2070-\\u218F\\u2C00-\\u2FEF\\u3001-\\uD7FF\\uF900-\\uFDCF\\uFDF0-\\uFFFD", i = new RegExp("^[" + n + "][" + n + "\\-.\\d\\u00B7\\u0300-\\u036F\\u203F-\\u2040]*$");
          function s(t2, e2) {
            const n2 = [];
            let i2 = e2.exec(t2);
            for (; i2; ) {
              const s2 = [];
              s2.startIndex = e2.lastIndex - i2[0].length;
              const r2 = i2.length;
              for (let t3 = 0; t3 < r2; t3++) s2.push(i2[t3]);
              n2.push(s2), i2 = e2.exec(t2);
            }
            return n2;
          }
          const r = function(t2) {
            return !(null == i.exec(t2));
          }, o = { allowBooleanAttributes: false, unpairedTags: [] };
          function a(t2, e2) {
            e2 = Object.assign({}, o, e2);
            const n2 = [];
            let i2 = false, s2 = false;
            "\uFEFF" === t2[0] && (t2 = t2.substr(1));
            for (let r2 = 0; r2 < t2.length; r2++) if ("<" === t2[r2] && "?" === t2[r2 + 1]) {
              if (r2 += 2, r2 = u(t2, r2), r2.err) return r2;
            } else {
              if ("<" !== t2[r2]) {
                if (l(t2[r2])) continue;
                return m("InvalidChar", "char '" + t2[r2] + "' is not expected.", N(t2, r2));
              }
              {
                let o2 = r2;
                if (r2++, "!" === t2[r2]) {
                  r2 = d(t2, r2);
                  continue;
                }
                {
                  let a2 = false;
                  "/" === t2[r2] && (a2 = true, r2++);
                  let h2 = "";
                  for (; r2 < t2.length && ">" !== t2[r2] && " " !== t2[r2] && "	" !== t2[r2] && "\n" !== t2[r2] && "\r" !== t2[r2]; r2++) h2 += t2[r2];
                  if (h2 = h2.trim(), "/" === h2[h2.length - 1] && (h2 = h2.substring(0, h2.length - 1), r2--), !b(h2)) {
                    let e3;
                    return e3 = 0 === h2.trim().length ? "Invalid space after '<'." : "Tag '" + h2 + "' is an invalid name.", m("InvalidTag", e3, N(t2, r2));
                  }
                  const p2 = c(t2, r2);
                  if (false === p2) return m("InvalidAttr", "Attributes for '" + h2 + "' have open quote.", N(t2, r2));
                  let f2 = p2.value;
                  if (r2 = p2.index, "/" === f2[f2.length - 1]) {
                    const n3 = r2 - f2.length;
                    f2 = f2.substring(0, f2.length - 1);
                    const s3 = g(f2, e2);
                    if (true !== s3) return m(s3.err.code, s3.err.msg, N(t2, n3 + s3.err.line));
                    i2 = true;
                  } else if (a2) {
                    if (!p2.tagClosed) return m("InvalidTag", "Closing tag '" + h2 + "' doesn't have proper closing.", N(t2, r2));
                    if (f2.trim().length > 0) return m("InvalidTag", "Closing tag '" + h2 + "' can't have attributes or invalid starting.", N(t2, o2));
                    if (0 === n2.length) return m("InvalidTag", "Closing tag '" + h2 + "' has not been opened.", N(t2, o2));
                    {
                      const e3 = n2.pop();
                      if (h2 !== e3.tagName) {
                        let n3 = N(t2, e3.tagStartPos);
                        return m("InvalidTag", "Expected closing tag '" + e3.tagName + "' (opened in line " + n3.line + ", col " + n3.col + ") instead of closing tag '" + h2 + "'.", N(t2, o2));
                      }
                      0 == n2.length && (s2 = true);
                    }
                  } else {
                    const a3 = g(f2, e2);
                    if (true !== a3) return m(a3.err.code, a3.err.msg, N(t2, r2 - f2.length + a3.err.line));
                    if (true === s2) return m("InvalidXml", "Multiple possible root nodes found.", N(t2, r2));
                    -1 !== e2.unpairedTags.indexOf(h2) || n2.push({ tagName: h2, tagStartPos: o2 }), i2 = true;
                  }
                  for (r2++; r2 < t2.length; r2++) if ("<" === t2[r2]) {
                    if ("!" === t2[r2 + 1]) {
                      r2++, r2 = d(t2, r2);
                      continue;
                    }
                    if ("?" !== t2[r2 + 1]) break;
                    if (r2 = u(t2, ++r2), r2.err) return r2;
                  } else if ("&" === t2[r2]) {
                    const e3 = x(t2, r2);
                    if (-1 == e3) return m("InvalidChar", "char '&' is not expected.", N(t2, r2));
                    r2 = e3;
                  } else if (true === s2 && !l(t2[r2])) return m("InvalidXml", "Extra text at the end", N(t2, r2));
                  "<" === t2[r2] && r2--;
                }
              }
            }
            return i2 ? 1 == n2.length ? m("InvalidTag", "Unclosed tag '" + n2[0].tagName + "'.", N(t2, n2[0].tagStartPos)) : !(n2.length > 0) || m("InvalidXml", "Invalid '" + JSON.stringify(n2.map((t3) => t3.tagName), null, 4).replace(/\r?\n/g, "") + "' found.", { line: 1, col: 1 }) : m("InvalidXml", "Start tag expected.", 1);
          }
          function l(t2) {
            return " " === t2 || "	" === t2 || "\n" === t2 || "\r" === t2;
          }
          function u(t2, e2) {
            const n2 = e2;
            for (; e2 < t2.length; e2++) if ("?" == t2[e2] || " " == t2[e2]) {
              const i2 = t2.substr(n2, e2 - n2);
              if (e2 > 5 && "xml" === i2) return m("InvalidXml", "XML declaration allowed only at the start of the document.", N(t2, e2));
              if ("?" == t2[e2] && ">" == t2[e2 + 1]) {
                e2++;
                break;
              }
              continue;
            }
            return e2;
          }
          function d(t2, e2) {
            if (t2.length > e2 + 5 && "-" === t2[e2 + 1] && "-" === t2[e2 + 2]) {
              for (e2 += 3; e2 < t2.length; e2++) if ("-" === t2[e2] && "-" === t2[e2 + 1] && ">" === t2[e2 + 2]) {
                e2 += 2;
                break;
              }
            } else if (t2.length > e2 + 8 && "D" === t2[e2 + 1] && "O" === t2[e2 + 2] && "C" === t2[e2 + 3] && "T" === t2[e2 + 4] && "Y" === t2[e2 + 5] && "P" === t2[e2 + 6] && "E" === t2[e2 + 7]) {
              let n2 = 1;
              for (e2 += 8; e2 < t2.length; e2++) if ("<" === t2[e2]) n2++;
              else if (">" === t2[e2] && (n2--, 0 === n2)) break;
            } else if (t2.length > e2 + 9 && "[" === t2[e2 + 1] && "C" === t2[e2 + 2] && "D" === t2[e2 + 3] && "A" === t2[e2 + 4] && "T" === t2[e2 + 5] && "A" === t2[e2 + 6] && "[" === t2[e2 + 7]) {
              for (e2 += 8; e2 < t2.length; e2++) if ("]" === t2[e2] && "]" === t2[e2 + 1] && ">" === t2[e2 + 2]) {
                e2 += 2;
                break;
              }
            }
            return e2;
          }
          const h = '"', p = "'";
          function c(t2, e2) {
            let n2 = "", i2 = "", s2 = false;
            for (; e2 < t2.length; e2++) {
              if (t2[e2] === h || t2[e2] === p) "" === i2 ? i2 = t2[e2] : i2 !== t2[e2] || (i2 = "");
              else if (">" === t2[e2] && "" === i2) {
                s2 = true;
                break;
              }
              n2 += t2[e2];
            }
            return "" === i2 && { value: n2, index: e2, tagClosed: s2 };
          }
          const f = new RegExp(`(\\s*)([^\\s=]+)(\\s*=)?(\\s*(['"])(([\\s\\S])*?)\\5)?`, "g");
          function g(t2, e2) {
            const n2 = s(t2, f), i2 = {};
            for (let t3 = 0; t3 < n2.length; t3++) {
              if (0 === n2[t3][1].length) return m("InvalidAttr", "Attribute '" + n2[t3][2] + "' has no space in starting.", y(n2[t3]));
              if (void 0 !== n2[t3][3] && void 0 === n2[t3][4]) return m("InvalidAttr", "Attribute '" + n2[t3][2] + "' is without value.", y(n2[t3]));
              if (void 0 === n2[t3][3] && !e2.allowBooleanAttributes) return m("InvalidAttr", "boolean attribute '" + n2[t3][2] + "' is not allowed.", y(n2[t3]));
              const s2 = n2[t3][2];
              if (!E(s2)) return m("InvalidAttr", "Attribute '" + s2 + "' is an invalid name.", y(n2[t3]));
              if (Object.prototype.hasOwnProperty.call(i2, s2)) return m("InvalidAttr", "Attribute '" + s2 + "' is repeated.", y(n2[t3]));
              i2[s2] = 1;
            }
            return true;
          }
          function x(t2, e2) {
            if (";" === t2[++e2]) return -1;
            if ("#" === t2[e2]) return (function(t3, e3) {
              let n3 = /\d/;
              for ("x" === t3[e3] && (e3++, n3 = /[\da-fA-F]/); e3 < t3.length; e3++) {
                if (";" === t3[e3]) return e3;
                if (!t3[e3].match(n3)) break;
              }
              return -1;
            })(t2, ++e2);
            let n2 = 0;
            for (; e2 < t2.length; e2++, n2++) if (!(t2[e2].match(/\w/) && n2 < 20)) {
              if (";" === t2[e2]) break;
              return -1;
            }
            return e2;
          }
          function m(t2, e2, n2) {
            return { err: { code: t2, msg: e2, line: n2.line || n2, col: n2.col } };
          }
          function E(t2) {
            return r(t2);
          }
          function b(t2) {
            return r(t2);
          }
          function N(t2, e2) {
            const n2 = t2.substring(0, e2).split(/\r?\n/);
            return { line: n2.length, col: n2[n2.length - 1].length + 1 };
          }
          function y(t2) {
            return t2.startIndex + t2[1].length;
          }
          const T = { preserveOrder: false, attributeNamePrefix: "@_", attributesGroupName: false, textNodeName: "#text", ignoreAttributes: true, removeNSPrefix: false, allowBooleanAttributes: false, parseTagValue: true, parseAttributeValue: false, trimValues: true, cdataPropName: false, numberParseOptions: { hex: true, leadingZeros: true, eNotation: true }, tagValueProcessor: function(t2, e2) {
            return e2;
          }, attributeValueProcessor: function(t2, e2) {
            return e2;
          }, stopNodes: [], alwaysCreateTextNode: false, isArray: () => false, commentPropName: false, unpairedTags: [], processEntities: true, htmlEntities: false, ignoreDeclaration: false, ignorePiTags: false, transformTagName: false, transformAttributeName: false, updateTag: function(t2, e2, n2) {
            return t2;
          }, captureMetaData: false, maxNestedTags: 100, strictReservedNames: true };
          function w(t2) {
            return "boolean" == typeof t2 ? { enabled: t2, maxEntitySize: 1e4, maxExpansionDepth: 10, maxTotalExpansions: 1e3, maxExpandedLength: 1e5, allowedTags: null, tagFilter: null } : "object" == typeof t2 && null !== t2 ? { enabled: false !== t2.enabled, maxEntitySize: t2.maxEntitySize ?? 1e4, maxExpansionDepth: t2.maxExpansionDepth ?? 10, maxTotalExpansions: t2.maxTotalExpansions ?? 1e3, maxExpandedLength: t2.maxExpandedLength ?? 1e5, allowedTags: t2.allowedTags ?? null, tagFilter: t2.tagFilter ?? null } : w(true);
          }
          const v = function(t2) {
            const e2 = Object.assign({}, T, t2);
            return e2.processEntities = w(e2.processEntities), e2;
          };
          let O;
          O = "function" != typeof Symbol ? "@@xmlMetadata" : /* @__PURE__ */ Symbol("XML Node Metadata");
          class I {
            constructor(t2) {
              this.tagname = t2, this.child = [], this[":@"] = /* @__PURE__ */ Object.create(null);
            }
            add(t2, e2) {
              "__proto__" === t2 && (t2 = "#__proto__"), this.child.push({ [t2]: e2 });
            }
            addChild(t2, e2) {
              "__proto__" === t2.tagname && (t2.tagname = "#__proto__"), t2[":@"] && Object.keys(t2[":@"]).length > 0 ? this.child.push({ [t2.tagname]: t2.child, ":@": t2[":@"] }) : this.child.push({ [t2.tagname]: t2.child }), void 0 !== e2 && (this.child[this.child.length - 1][O] = { startIndex: e2 });
            }
            static getMetaDataSymbol() {
              return O;
            }
          }
          class P {
            constructor(t2) {
              this.suppressValidationErr = !t2, this.options = t2;
            }
            readDocType(t2, e2) {
              const n2 = /* @__PURE__ */ Object.create(null);
              if ("O" !== t2[e2 + 3] || "C" !== t2[e2 + 4] || "T" !== t2[e2 + 5] || "Y" !== t2[e2 + 6] || "P" !== t2[e2 + 7] || "E" !== t2[e2 + 8]) throw new Error("Invalid Tag instead of DOCTYPE");
              {
                e2 += 9;
                let i2 = 1, s2 = false, r2 = false, o2 = "";
                for (; e2 < t2.length; e2++) if ("<" !== t2[e2] || r2) if (">" === t2[e2]) {
                  if (r2 ? "-" === t2[e2 - 1] && "-" === t2[e2 - 2] && (r2 = false, i2--) : i2--, 0 === i2) break;
                } else "[" === t2[e2] ? s2 = true : o2 += t2[e2];
                else {
                  if (s2 && S(t2, "!ENTITY", e2)) {
                    let i3, s3;
                    if (e2 += 7, [i3, s3, e2] = this.readEntityExp(t2, e2 + 1, this.suppressValidationErr), -1 === s3.indexOf("&")) {
                      const t3 = i3.replace(/[.\-+*:]/g, "\\.");
                      n2[i3] = { regx: RegExp(`&${t3};`, "g"), val: s3 };
                    }
                  } else if (s2 && S(t2, "!ELEMENT", e2)) {
                    e2 += 8;
                    const { index: n3 } = this.readElementExp(t2, e2 + 1);
                    e2 = n3;
                  } else if (s2 && S(t2, "!ATTLIST", e2)) e2 += 8;
                  else if (s2 && S(t2, "!NOTATION", e2)) {
                    e2 += 9;
                    const { index: n3 } = this.readNotationExp(t2, e2 + 1, this.suppressValidationErr);
                    e2 = n3;
                  } else {
                    if (!S(t2, "!--", e2)) throw new Error("Invalid DOCTYPE");
                    r2 = true;
                  }
                  i2++, o2 = "";
                }
                if (0 !== i2) throw new Error("Unclosed DOCTYPE");
              }
              return { entities: n2, i: e2 };
            }
            readEntityExp(t2, e2) {
              e2 = A(t2, e2);
              let n2 = "";
              for (; e2 < t2.length && !/\s/.test(t2[e2]) && '"' !== t2[e2] && "'" !== t2[e2]; ) n2 += t2[e2], e2++;
              if (C(n2), e2 = A(t2, e2), !this.suppressValidationErr) {
                if ("SYSTEM" === t2.substring(e2, e2 + 6).toUpperCase()) throw new Error("External entities are not supported");
                if ("%" === t2[e2]) throw new Error("Parameter entities are not supported");
              }
              let i2 = "";
              if ([e2, i2] = this.readIdentifierVal(t2, e2, "entity"), false !== this.options.enabled && this.options.maxEntitySize && i2.length > this.options.maxEntitySize) throw new Error(`Entity "${n2}" size (${i2.length}) exceeds maximum allowed size (${this.options.maxEntitySize})`);
              return [n2, i2, --e2];
            }
            readNotationExp(t2, e2) {
              e2 = A(t2, e2);
              let n2 = "";
              for (; e2 < t2.length && !/\s/.test(t2[e2]); ) n2 += t2[e2], e2++;
              !this.suppressValidationErr && C(n2), e2 = A(t2, e2);
              const i2 = t2.substring(e2, e2 + 6).toUpperCase();
              if (!this.suppressValidationErr && "SYSTEM" !== i2 && "PUBLIC" !== i2) throw new Error(`Expected SYSTEM or PUBLIC, found "${i2}"`);
              e2 += i2.length, e2 = A(t2, e2);
              let s2 = null, r2 = null;
              if ("PUBLIC" === i2) [e2, s2] = this.readIdentifierVal(t2, e2, "publicIdentifier"), '"' !== t2[e2 = A(t2, e2)] && "'" !== t2[e2] || ([e2, r2] = this.readIdentifierVal(t2, e2, "systemIdentifier"));
              else if ("SYSTEM" === i2 && ([e2, r2] = this.readIdentifierVal(t2, e2, "systemIdentifier"), !this.suppressValidationErr && !r2)) throw new Error("Missing mandatory system identifier for SYSTEM notation");
              return { notationName: n2, publicIdentifier: s2, systemIdentifier: r2, index: --e2 };
            }
            readIdentifierVal(t2, e2, n2) {
              let i2 = "";
              const s2 = t2[e2];
              if ('"' !== s2 && "'" !== s2) throw new Error(`Expected quoted string, found "${s2}"`);
              for (e2++; e2 < t2.length && t2[e2] !== s2; ) i2 += t2[e2], e2++;
              if (t2[e2] !== s2) throw new Error(`Unterminated ${n2} value`);
              return [++e2, i2];
            }
            readElementExp(t2, e2) {
              e2 = A(t2, e2);
              let n2 = "";
              for (; e2 < t2.length && !/\s/.test(t2[e2]); ) n2 += t2[e2], e2++;
              if (!this.suppressValidationErr && !r(n2)) throw new Error(`Invalid element name: "${n2}"`);
              let i2 = "";
              if ("E" === t2[e2 = A(t2, e2)] && S(t2, "MPTY", e2)) e2 += 4;
              else if ("A" === t2[e2] && S(t2, "NY", e2)) e2 += 2;
              else if ("(" === t2[e2]) {
                for (e2++; e2 < t2.length && ")" !== t2[e2]; ) i2 += t2[e2], e2++;
                if (")" !== t2[e2]) throw new Error("Unterminated content model");
              } else if (!this.suppressValidationErr) throw new Error(`Invalid Element Expression, found "${t2[e2]}"`);
              return { elementName: n2, contentModel: i2.trim(), index: e2 };
            }
            readAttlistExp(t2, e2) {
              e2 = A(t2, e2);
              let n2 = "";
              for (; e2 < t2.length && !/\s/.test(t2[e2]); ) n2 += t2[e2], e2++;
              C(n2), e2 = A(t2, e2);
              let i2 = "";
              for (; e2 < t2.length && !/\s/.test(t2[e2]); ) i2 += t2[e2], e2++;
              if (!C(i2)) throw new Error(`Invalid attribute name: "${i2}"`);
              e2 = A(t2, e2);
              let s2 = "";
              if ("NOTATION" === t2.substring(e2, e2 + 8).toUpperCase()) {
                if (s2 = "NOTATION", "(" !== t2[e2 = A(t2, e2 += 8)]) throw new Error(`Expected '(', found "${t2[e2]}"`);
                e2++;
                let n3 = [];
                for (; e2 < t2.length && ")" !== t2[e2]; ) {
                  let i3 = "";
                  for (; e2 < t2.length && "|" !== t2[e2] && ")" !== t2[e2]; ) i3 += t2[e2], e2++;
                  if (i3 = i3.trim(), !C(i3)) throw new Error(`Invalid notation name: "${i3}"`);
                  n3.push(i3), "|" === t2[e2] && (e2++, e2 = A(t2, e2));
                }
                if (")" !== t2[e2]) throw new Error("Unterminated list of notations");
                e2++, s2 += " (" + n3.join("|") + ")";
              } else {
                for (; e2 < t2.length && !/\s/.test(t2[e2]); ) s2 += t2[e2], e2++;
                const n3 = ["CDATA", "ID", "IDREF", "IDREFS", "ENTITY", "ENTITIES", "NMTOKEN", "NMTOKENS"];
                if (!this.suppressValidationErr && !n3.includes(s2.toUpperCase())) throw new Error(`Invalid attribute type: "${s2}"`);
              }
              e2 = A(t2, e2);
              let r2 = "";
              return "#REQUIRED" === t2.substring(e2, e2 + 8).toUpperCase() ? (r2 = "#REQUIRED", e2 += 8) : "#IMPLIED" === t2.substring(e2, e2 + 7).toUpperCase() ? (r2 = "#IMPLIED", e2 += 7) : [e2, r2] = this.readIdentifierVal(t2, e2, "ATTLIST"), { elementName: n2, attributeName: i2, attributeType: s2, defaultValue: r2, index: e2 };
            }
          }
          const A = (t2, e2) => {
            for (; e2 < t2.length && /\s/.test(t2[e2]); ) e2++;
            return e2;
          };
          function S(t2, e2, n2) {
            for (let i2 = 0; i2 < e2.length; i2++) if (e2[i2] !== t2[n2 + i2 + 1]) return false;
            return true;
          }
          function C(t2) {
            if (r(t2)) return t2;
            throw new Error(`Invalid entity name ${t2}`);
          }
          const $ = /^[-+]?0x[a-fA-F0-9]+$/, V = /^([\-\+])?(0*)([0-9]*(\.[0-9]*)?)$/, D = { hex: true, leadingZeros: true, decimalPoint: ".", eNotation: true };
          const j = /^([-+])?(0*)(\d*(\.\d*)?[eE][-\+]?\d+)$/;
          class L {
            constructor(t2) {
              var e2;
              if (this.options = t2, this.currentNode = null, this.tagsNodeStack = [], this.docTypeEntities = {}, this.lastEntities = { apos: { regex: /&(apos|#39|#x27);/g, val: "'" }, gt: { regex: /&(gt|#62|#x3E);/g, val: ">" }, lt: { regex: /&(lt|#60|#x3C);/g, val: "<" }, quot: { regex: /&(quot|#34|#x22);/g, val: '"' } }, this.ampEntity = { regex: /&(amp|#38|#x26);/g, val: "&" }, this.htmlEntities = { space: { regex: /&(nbsp|#160);/g, val: " " }, cent: { regex: /&(cent|#162);/g, val: "\xA2" }, pound: { regex: /&(pound|#163);/g, val: "\xA3" }, yen: { regex: /&(yen|#165);/g, val: "\xA5" }, euro: { regex: /&(euro|#8364);/g, val: "\u20AC" }, copyright: { regex: /&(copy|#169);/g, val: "\xA9" }, reg: { regex: /&(reg|#174);/g, val: "\xAE" }, inr: { regex: /&(inr|#8377);/g, val: "\u20B9" }, num_dec: { regex: /&#([0-9]{1,7});/g, val: (t3, e3) => K(e3, 10, "&#") }, num_hex: { regex: /&#x([0-9a-fA-F]{1,6});/g, val: (t3, e3) => K(e3, 16, "&#x") } }, this.addExternalEntities = F, this.parseXml = R, this.parseTextData = M, this.resolveNameSpace = k, this.buildAttributesMap = U, this.isItStopNode = X, this.replaceEntitiesValue = Y, this.readStopNodeData = q, this.saveTextToParentTag = G, this.addChild = B, this.ignoreAttributesFn = "function" == typeof (e2 = this.options.ignoreAttributes) ? e2 : Array.isArray(e2) ? (t3) => {
                for (const n2 of e2) {
                  if ("string" == typeof n2 && t3 === n2) return true;
                  if (n2 instanceof RegExp && n2.test(t3)) return true;
                }
              } : () => false, this.entityExpansionCount = 0, this.currentExpandedLength = 0, this.options.stopNodes && this.options.stopNodes.length > 0) {
                this.stopNodesExact = /* @__PURE__ */ new Set(), this.stopNodesWildcard = /* @__PURE__ */ new Set();
                for (let t3 = 0; t3 < this.options.stopNodes.length; t3++) {
                  const e3 = this.options.stopNodes[t3];
                  "string" == typeof e3 && (e3.startsWith("*.") ? this.stopNodesWildcard.add(e3.substring(2)) : this.stopNodesExact.add(e3));
                }
              }
            }
          }
          function F(t2) {
            const e2 = Object.keys(t2);
            for (let n2 = 0; n2 < e2.length; n2++) {
              const i2 = e2[n2], s2 = i2.replace(/[.\-+*:]/g, "\\.");
              this.lastEntities[i2] = { regex: new RegExp("&" + s2 + ";", "g"), val: t2[i2] };
            }
          }
          function M(t2, e2, n2, i2, s2, r2, o2) {
            if (void 0 !== t2 && (this.options.trimValues && !i2 && (t2 = t2.trim()), t2.length > 0)) {
              o2 || (t2 = this.replaceEntitiesValue(t2, e2, n2));
              const i3 = this.options.tagValueProcessor(e2, t2, n2, s2, r2);
              return null == i3 ? t2 : typeof i3 != typeof t2 || i3 !== t2 ? i3 : this.options.trimValues || t2.trim() === t2 ? Z(t2, this.options.parseTagValue, this.options.numberParseOptions) : t2;
            }
          }
          function k(t2) {
            if (this.options.removeNSPrefix) {
              const e2 = t2.split(":"), n2 = "/" === t2.charAt(0) ? "/" : "";
              if ("xmlns" === e2[0]) return "";
              2 === e2.length && (t2 = n2 + e2[1]);
            }
            return t2;
          }
          const _ = new RegExp(`([^\\s=]+)\\s*(=\\s*(['"])([\\s\\S]*?)\\3)?`, "gm");
          function U(t2, e2, n2) {
            if (true !== this.options.ignoreAttributes && "string" == typeof t2) {
              const i2 = s(t2, _), r2 = i2.length, o2 = {};
              for (let t3 = 0; t3 < r2; t3++) {
                const s2 = this.resolveNameSpace(i2[t3][1]);
                if (this.ignoreAttributesFn(s2, e2)) continue;
                let r3 = i2[t3][4], a2 = this.options.attributeNamePrefix + s2;
                if (s2.length) if (this.options.transformAttributeName && (a2 = this.options.transformAttributeName(a2)), "__proto__" === a2 && (a2 = "#__proto__"), void 0 !== r3) {
                  this.options.trimValues && (r3 = r3.trim()), r3 = this.replaceEntitiesValue(r3, n2, e2);
                  const t4 = this.options.attributeValueProcessor(s2, r3, e2);
                  o2[a2] = null == t4 ? r3 : typeof t4 != typeof r3 || t4 !== r3 ? t4 : Z(r3, this.options.parseAttributeValue, this.options.numberParseOptions);
                } else this.options.allowBooleanAttributes && (o2[a2] = true);
              }
              if (!Object.keys(o2).length) return;
              if (this.options.attributesGroupName) {
                const t3 = {};
                return t3[this.options.attributesGroupName] = o2, t3;
              }
              return o2;
            }
          }
          const R = function(t2) {
            t2 = t2.replace(/\r\n?/g, "\n");
            const e2 = new I("!xml");
            let n2 = e2, i2 = "", s2 = "";
            this.entityExpansionCount = 0, this.currentExpandedLength = 0;
            const r2 = new P(this.options.processEntities);
            for (let o2 = 0; o2 < t2.length; o2++) if ("<" === t2[o2]) if ("/" === t2[o2 + 1]) {
              const e3 = z(t2, ">", o2, "Closing Tag is not closed.");
              let r3 = t2.substring(o2 + 2, e3).trim();
              if (this.options.removeNSPrefix) {
                const t3 = r3.indexOf(":");
                -1 !== t3 && (r3 = r3.substr(t3 + 1));
              }
              this.options.transformTagName && (r3 = this.options.transformTagName(r3)), n2 && (i2 = this.saveTextToParentTag(i2, n2, s2));
              const a2 = s2.substring(s2.lastIndexOf(".") + 1);
              if (r3 && -1 !== this.options.unpairedTags.indexOf(r3)) throw new Error(`Unpaired tag can not be used as closing tag: </${r3}>`);
              let l2 = 0;
              a2 && -1 !== this.options.unpairedTags.indexOf(a2) ? (l2 = s2.lastIndexOf(".", s2.lastIndexOf(".") - 1), this.tagsNodeStack.pop()) : l2 = s2.lastIndexOf("."), s2 = s2.substring(0, l2), n2 = this.tagsNodeStack.pop(), i2 = "", o2 = e3;
            } else if ("?" === t2[o2 + 1]) {
              let e3 = W(t2, o2, false, "?>");
              if (!e3) throw new Error("Pi Tag is not closed.");
              if (i2 = this.saveTextToParentTag(i2, n2, s2), this.options.ignoreDeclaration && "?xml" === e3.tagName || this.options.ignorePiTags) ;
              else {
                const t3 = new I(e3.tagName);
                t3.add(this.options.textNodeName, ""), e3.tagName !== e3.tagExp && e3.attrExpPresent && (t3[":@"] = this.buildAttributesMap(e3.tagExp, s2, e3.tagName)), this.addChild(n2, t3, s2, o2);
              }
              o2 = e3.closeIndex + 1;
            } else if ("!--" === t2.substr(o2 + 1, 3)) {
              const e3 = z(t2, "-->", o2 + 4, "Comment is not closed.");
              if (this.options.commentPropName) {
                const r3 = t2.substring(o2 + 4, e3 - 2);
                i2 = this.saveTextToParentTag(i2, n2, s2), n2.add(this.options.commentPropName, [{ [this.options.textNodeName]: r3 }]);
              }
              o2 = e3;
            } else if ("!D" === t2.substr(o2 + 1, 2)) {
              const e3 = r2.readDocType(t2, o2);
              this.docTypeEntities = e3.entities, o2 = e3.i;
            } else if ("![" === t2.substr(o2 + 1, 2)) {
              const e3 = z(t2, "]]>", o2, "CDATA is not closed.") - 2, r3 = t2.substring(o2 + 9, e3);
              i2 = this.saveTextToParentTag(i2, n2, s2);
              let a2 = this.parseTextData(r3, n2.tagname, s2, true, false, true, true);
              null == a2 && (a2 = ""), this.options.cdataPropName ? n2.add(this.options.cdataPropName, [{ [this.options.textNodeName]: r3 }]) : n2.add(this.options.textNodeName, a2), o2 = e3 + 2;
            } else {
              let r3 = W(t2, o2, this.options.removeNSPrefix), a2 = r3.tagName;
              const l2 = r3.rawTagName;
              let u2 = r3.tagExp, d2 = r3.attrExpPresent, h2 = r3.closeIndex;
              if (this.options.transformTagName) {
                const t3 = this.options.transformTagName(a2);
                u2 === a2 && (u2 = t3), a2 = t3;
              }
              if (this.options.strictReservedNames && (a2 === this.options.commentPropName || a2 === this.options.cdataPropName)) throw new Error(`Invalid tag name: ${a2}`);
              n2 && i2 && "!xml" !== n2.tagname && (i2 = this.saveTextToParentTag(i2, n2, s2, false));
              const p2 = n2;
              p2 && -1 !== this.options.unpairedTags.indexOf(p2.tagname) && (n2 = this.tagsNodeStack.pop(), s2 = s2.substring(0, s2.lastIndexOf("."))), a2 !== e2.tagname && (s2 += s2 ? "." + a2 : a2);
              const c2 = o2;
              if (this.isItStopNode(this.stopNodesExact, this.stopNodesWildcard, s2, a2)) {
                let e3 = "";
                if (u2.length > 0 && u2.lastIndexOf("/") === u2.length - 1) "/" === a2[a2.length - 1] ? (a2 = a2.substr(0, a2.length - 1), s2 = s2.substr(0, s2.length - 1), u2 = a2) : u2 = u2.substr(0, u2.length - 1), o2 = r3.closeIndex;
                else if (-1 !== this.options.unpairedTags.indexOf(a2)) o2 = r3.closeIndex;
                else {
                  const n3 = this.readStopNodeData(t2, l2, h2 + 1);
                  if (!n3) throw new Error(`Unexpected end of ${l2}`);
                  o2 = n3.i, e3 = n3.tagContent;
                }
                const i3 = new I(a2);
                a2 !== u2 && d2 && (i3[":@"] = this.buildAttributesMap(u2, s2, a2)), e3 && (e3 = this.parseTextData(e3, a2, s2, true, d2, true, true)), s2 = s2.substr(0, s2.lastIndexOf(".")), i3.add(this.options.textNodeName, e3), this.addChild(n2, i3, s2, c2);
              } else {
                if (u2.length > 0 && u2.lastIndexOf("/") === u2.length - 1) {
                  if ("/" === a2[a2.length - 1] ? (a2 = a2.substr(0, a2.length - 1), s2 = s2.substr(0, s2.length - 1), u2 = a2) : u2 = u2.substr(0, u2.length - 1), this.options.transformTagName) {
                    const t4 = this.options.transformTagName(a2);
                    u2 === a2 && (u2 = t4), a2 = t4;
                  }
                  const t3 = new I(a2);
                  a2 !== u2 && d2 && (t3[":@"] = this.buildAttributesMap(u2, s2, a2)), this.addChild(n2, t3, s2, c2), s2 = s2.substr(0, s2.lastIndexOf("."));
                } else {
                  if (-1 !== this.options.unpairedTags.indexOf(a2)) {
                    const t3 = new I(a2);
                    a2 !== u2 && d2 && (t3[":@"] = this.buildAttributesMap(u2, s2)), this.addChild(n2, t3, s2, c2), s2 = s2.substr(0, s2.lastIndexOf(".")), o2 = r3.closeIndex;
                    continue;
                  }
                  {
                    const t3 = new I(a2);
                    if (this.tagsNodeStack.length > this.options.maxNestedTags) throw new Error("Maximum nested tags exceeded");
                    this.tagsNodeStack.push(n2), a2 !== u2 && d2 && (t3[":@"] = this.buildAttributesMap(u2, s2, a2)), this.addChild(n2, t3, s2, c2), n2 = t3;
                  }
                }
                i2 = "", o2 = h2;
              }
            }
            else i2 += t2[o2];
            return e2.child;
          };
          function B(t2, e2, n2, i2) {
            this.options.captureMetaData || (i2 = void 0);
            const s2 = this.options.updateTag(e2.tagname, n2, e2[":@"]);
            false === s2 || ("string" == typeof s2 ? (e2.tagname = s2, t2.addChild(e2, i2)) : t2.addChild(e2, i2));
          }
          const Y = function(t2, e2, n2) {
            if (-1 === t2.indexOf("&")) return t2;
            const i2 = this.options.processEntities;
            if (!i2.enabled) return t2;
            if (i2.allowedTags && !i2.allowedTags.includes(e2)) return t2;
            if (i2.tagFilter && !i2.tagFilter(e2, n2)) return t2;
            for (let e3 in this.docTypeEntities) {
              const n3 = this.docTypeEntities[e3], s2 = t2.match(n3.regx);
              if (s2) {
                if (this.entityExpansionCount += s2.length, i2.maxTotalExpansions && this.entityExpansionCount > i2.maxTotalExpansions) throw new Error(`Entity expansion limit exceeded: ${this.entityExpansionCount} > ${i2.maxTotalExpansions}`);
                const e4 = t2.length;
                if (t2 = t2.replace(n3.regx, n3.val), i2.maxExpandedLength && (this.currentExpandedLength += t2.length - e4, this.currentExpandedLength > i2.maxExpandedLength)) throw new Error(`Total expanded content size exceeded: ${this.currentExpandedLength} > ${i2.maxExpandedLength}`);
              }
            }
            if (-1 === t2.indexOf("&")) return t2;
            for (let e3 in this.lastEntities) {
              const n3 = this.lastEntities[e3];
              t2 = t2.replace(n3.regex, n3.val);
            }
            if (-1 === t2.indexOf("&")) return t2;
            if (this.options.htmlEntities) for (let e3 in this.htmlEntities) {
              const n3 = this.htmlEntities[e3];
              t2 = t2.replace(n3.regex, n3.val);
            }
            return t2.replace(this.ampEntity.regex, this.ampEntity.val);
          };
          function G(t2, e2, n2, i2) {
            return t2 && (void 0 === i2 && (i2 = 0 === e2.child.length), void 0 !== (t2 = this.parseTextData(t2, e2.tagname, n2, false, !!e2[":@"] && 0 !== Object.keys(e2[":@"]).length, i2)) && "" !== t2 && e2.add(this.options.textNodeName, t2), t2 = ""), t2;
          }
          function X(t2, e2, n2, i2) {
            return !(!e2 || !e2.has(i2)) || !(!t2 || !t2.has(n2));
          }
          function z(t2, e2, n2, i2) {
            const s2 = t2.indexOf(e2, n2);
            if (-1 === s2) throw new Error(i2);
            return s2 + e2.length - 1;
          }
          function W(t2, e2, n2, i2 = ">") {
            const s2 = (function(t3, e3, n3 = ">") {
              let i3, s3 = "";
              for (let r3 = e3; r3 < t3.length; r3++) {
                let e4 = t3[r3];
                if (i3) e4 === i3 && (i3 = "");
                else if ('"' === e4 || "'" === e4) i3 = e4;
                else if (e4 === n3[0]) {
                  if (!n3[1]) return { data: s3, index: r3 };
                  if (t3[r3 + 1] === n3[1]) return { data: s3, index: r3 };
                } else "	" === e4 && (e4 = " ");
                s3 += e4;
              }
            })(t2, e2 + 1, i2);
            if (!s2) return;
            let r2 = s2.data;
            const o2 = s2.index, a2 = r2.search(/\s/);
            let l2 = r2, u2 = true;
            -1 !== a2 && (l2 = r2.substring(0, a2), r2 = r2.substring(a2 + 1).trimStart());
            const d2 = l2;
            if (n2) {
              const t3 = l2.indexOf(":");
              -1 !== t3 && (l2 = l2.substr(t3 + 1), u2 = l2 !== s2.data.substr(t3 + 1));
            }
            return { tagName: l2, tagExp: r2, closeIndex: o2, attrExpPresent: u2, rawTagName: d2 };
          }
          function q(t2, e2, n2) {
            const i2 = n2;
            let s2 = 1;
            for (; n2 < t2.length; n2++) if ("<" === t2[n2]) if ("/" === t2[n2 + 1]) {
              const r2 = z(t2, ">", n2, `${e2} is not closed`);
              if (t2.substring(n2 + 2, r2).trim() === e2 && (s2--, 0 === s2)) return { tagContent: t2.substring(i2, n2), i: r2 };
              n2 = r2;
            } else if ("?" === t2[n2 + 1]) n2 = z(t2, "?>", n2 + 1, "StopNode is not closed.");
            else if ("!--" === t2.substr(n2 + 1, 3)) n2 = z(t2, "-->", n2 + 3, "StopNode is not closed.");
            else if ("![" === t2.substr(n2 + 1, 2)) n2 = z(t2, "]]>", n2, "StopNode is not closed.") - 2;
            else {
              const i3 = W(t2, n2, ">");
              i3 && ((i3 && i3.tagName) === e2 && "/" !== i3.tagExp[i3.tagExp.length - 1] && s2++, n2 = i3.closeIndex);
            }
          }
          function Z(t2, e2, n2) {
            if (e2 && "string" == typeof t2) {
              const e3 = t2.trim();
              return "true" === e3 || "false" !== e3 && (function(t3, e4 = {}) {
                if (e4 = Object.assign({}, D, e4), !t3 || "string" != typeof t3) return t3;
                let n3 = t3.trim();
                if (void 0 !== e4.skipLike && e4.skipLike.test(n3)) return t3;
                if ("0" === t3) return 0;
                if (e4.hex && $.test(n3)) return (function(t4) {
                  if (parseInt) return parseInt(t4, 16);
                  if (Number.parseInt) return Number.parseInt(t4, 16);
                  if (window && window.parseInt) return window.parseInt(t4, 16);
                  throw new Error("parseInt, Number.parseInt, window.parseInt are not supported");
                })(n3);
                if (n3.includes("e") || n3.includes("E")) return (function(t4, e5, n4) {
                  if (!n4.eNotation) return t4;
                  const i3 = e5.match(j);
                  if (i3) {
                    let s2 = i3[1] || "";
                    const r2 = -1 === i3[3].indexOf("e") ? "E" : "e", o2 = i3[2], a2 = s2 ? t4[o2.length + 1] === r2 : t4[o2.length] === r2;
                    return o2.length > 1 && a2 ? t4 : 1 !== o2.length || !i3[3].startsWith(`.${r2}`) && i3[3][0] !== r2 ? n4.leadingZeros && !a2 ? (e5 = (i3[1] || "") + i3[3], Number(e5)) : t4 : Number(e5);
                  }
                  return t4;
                })(t3, n3, e4);
                {
                  const s2 = V.exec(n3);
                  if (s2) {
                    const r2 = s2[1] || "", o2 = s2[2];
                    let a2 = (i2 = s2[3]) && -1 !== i2.indexOf(".") ? ("." === (i2 = i2.replace(/0+$/, "")) ? i2 = "0" : "." === i2[0] ? i2 = "0" + i2 : "." === i2[i2.length - 1] && (i2 = i2.substring(0, i2.length - 1)), i2) : i2;
                    const l2 = r2 ? "." === t3[o2.length + 1] : "." === t3[o2.length];
                    if (!e4.leadingZeros && (o2.length > 1 || 1 === o2.length && !l2)) return t3;
                    {
                      const i3 = Number(n3), s3 = String(i3);
                      if (0 === i3) return i3;
                      if (-1 !== s3.search(/[eE]/)) return e4.eNotation ? i3 : t3;
                      if (-1 !== n3.indexOf(".")) return "0" === s3 || s3 === a2 || s3 === `${r2}${a2}` ? i3 : t3;
                      let l3 = o2 ? a2 : n3;
                      return o2 ? l3 === s3 || r2 + l3 === s3 ? i3 : t3 : l3 === s3 || l3 === r2 + s3 ? i3 : t3;
                    }
                  }
                  return t3;
                }
                var i2;
              })(t2, n2);
            }
            return void 0 !== t2 ? t2 : "";
          }
          function K(t2, e2, n2) {
            const i2 = Number.parseInt(t2, e2);
            return i2 >= 0 && i2 <= 1114111 ? String.fromCodePoint(i2) : n2 + t2 + ";";
          }
          const Q = I.getMetaDataSymbol();
          function J(t2, e2) {
            return H(t2, e2);
          }
          function H(t2, e2, n2) {
            let i2;
            const s2 = {};
            for (let r2 = 0; r2 < t2.length; r2++) {
              const o2 = t2[r2], a2 = tt(o2);
              let l2 = "";
              if (l2 = void 0 === n2 ? a2 : n2 + "." + a2, a2 === e2.textNodeName) void 0 === i2 ? i2 = o2[a2] : i2 += "" + o2[a2];
              else {
                if (void 0 === a2) continue;
                if (o2[a2]) {
                  let t3 = H(o2[a2], e2, l2);
                  const n3 = nt(t3, e2);
                  o2[":@"] ? et(t3, o2[":@"], l2, e2) : 1 !== Object.keys(t3).length || void 0 === t3[e2.textNodeName] || e2.alwaysCreateTextNode ? 0 === Object.keys(t3).length && (e2.alwaysCreateTextNode ? t3[e2.textNodeName] = "" : t3 = "") : t3 = t3[e2.textNodeName], void 0 !== o2[Q] && "object" == typeof t3 && null !== t3 && (t3[Q] = o2[Q]), void 0 !== s2[a2] && Object.prototype.hasOwnProperty.call(s2, a2) ? (Array.isArray(s2[a2]) || (s2[a2] = [s2[a2]]), s2[a2].push(t3)) : e2.isArray(a2, l2, n3) ? s2[a2] = [t3] : s2[a2] = t3;
                }
              }
            }
            return "string" == typeof i2 ? i2.length > 0 && (s2[e2.textNodeName] = i2) : void 0 !== i2 && (s2[e2.textNodeName] = i2), s2;
          }
          function tt(t2) {
            const e2 = Object.keys(t2);
            for (let t3 = 0; t3 < e2.length; t3++) {
              const n2 = e2[t3];
              if (":@" !== n2) return n2;
            }
          }
          function et(t2, e2, n2, i2) {
            if (e2) {
              const s2 = Object.keys(e2), r2 = s2.length;
              for (let o2 = 0; o2 < r2; o2++) {
                const r3 = s2[o2];
                i2.isArray(r3, n2 + "." + r3, true, true) ? t2[r3] = [e2[r3]] : t2[r3] = e2[r3];
              }
            }
          }
          function nt(t2, e2) {
            const { textNodeName: n2 } = e2, i2 = Object.keys(t2).length;
            return 0 === i2 || !(1 !== i2 || !t2[n2] && "boolean" != typeof t2[n2] && 0 !== t2[n2]);
          }
          class it {
            constructor(t2) {
              this.externalEntities = {}, this.options = v(t2);
            }
            parse(t2, e2) {
              if ("string" != typeof t2 && t2.toString) t2 = t2.toString();
              else if ("string" != typeof t2) throw new Error("XML data is accepted in String or Bytes[] form.");
              if (e2) {
                true === e2 && (e2 = {});
                const n3 = a(t2, e2);
                if (true !== n3) throw Error(`${n3.err.msg}:${n3.err.line}:${n3.err.col}`);
              }
              const n2 = new L(this.options);
              n2.addExternalEntities(this.externalEntities);
              const i2 = n2.parseXml(t2);
              return this.options.preserveOrder || void 0 === i2 ? i2 : J(i2, this.options);
            }
            addEntity(t2, e2) {
              if (-1 !== e2.indexOf("&")) throw new Error("Entity value can't have '&'");
              if (-1 !== t2.indexOf("&") || -1 !== t2.indexOf(";")) throw new Error("An entity must be set without '&' and ';'. Eg. use '#xD' for '&#xD;'");
              if ("&" === e2) throw new Error("An entity with value '&' is not permitted");
              this.externalEntities[t2] = e2;
            }
            static getMetaDataSymbol() {
              return I.getMetaDataSymbol();
            }
          }
          function st(t2, e2) {
            let n2 = "";
            return e2.format && e2.indentBy.length > 0 && (n2 = "\n"), rt(t2, e2, "", n2);
          }
          function rt(t2, e2, n2, i2) {
            let s2 = "", r2 = false;
            if (!Array.isArray(t2)) {
              if (null != t2) {
                let n3 = t2.toString();
                return n3 = ut(n3, e2), n3;
              }
              return "";
            }
            for (let o2 = 0; o2 < t2.length; o2++) {
              const a2 = t2[o2], l2 = ot(a2);
              if (void 0 === l2) continue;
              let u2 = "";
              if (u2 = 0 === n2.length ? l2 : `${n2}.${l2}`, l2 === e2.textNodeName) {
                let t3 = a2[l2];
                lt(u2, e2) || (t3 = e2.tagValueProcessor(l2, t3), t3 = ut(t3, e2)), r2 && (s2 += i2), s2 += t3, r2 = false;
                continue;
              }
              if (l2 === e2.cdataPropName) {
                r2 && (s2 += i2), s2 += `<![CDATA[${a2[l2][0][e2.textNodeName]}]]>`, r2 = false;
                continue;
              }
              if (l2 === e2.commentPropName) {
                s2 += i2 + `<!--${a2[l2][0][e2.textNodeName]}-->`, r2 = true;
                continue;
              }
              if ("?" === l2[0]) {
                const t3 = at(a2[":@"], e2), n3 = "?xml" === l2 ? "" : i2;
                let o3 = a2[l2][0][e2.textNodeName];
                o3 = 0 !== o3.length ? " " + o3 : "", s2 += n3 + `<${l2}${o3}${t3}?>`, r2 = true;
                continue;
              }
              let d2 = i2;
              "" !== d2 && (d2 += e2.indentBy);
              const h2 = i2 + `<${l2}${at(a2[":@"], e2)}`, p2 = rt(a2[l2], e2, u2, d2);
              -1 !== e2.unpairedTags.indexOf(l2) ? e2.suppressUnpairedNode ? s2 += h2 + ">" : s2 += h2 + "/>" : p2 && 0 !== p2.length || !e2.suppressEmptyNode ? p2 && p2.endsWith(">") ? s2 += h2 + `>${p2}${i2}</${l2}>` : (s2 += h2 + ">", p2 && "" !== i2 && (p2.includes("/>") || p2.includes("</")) ? s2 += i2 + e2.indentBy + p2 + i2 : s2 += p2, s2 += `</${l2}>`) : s2 += h2 + "/>", r2 = true;
            }
            return s2;
          }
          function ot(t2) {
            const e2 = Object.keys(t2);
            for (let n2 = 0; n2 < e2.length; n2++) {
              const i2 = e2[n2];
              if (Object.prototype.hasOwnProperty.call(t2, i2) && ":@" !== i2) return i2;
            }
          }
          function at(t2, e2) {
            let n2 = "";
            if (t2 && !e2.ignoreAttributes) for (let i2 in t2) {
              if (!Object.prototype.hasOwnProperty.call(t2, i2)) continue;
              let s2 = e2.attributeValueProcessor(i2, t2[i2]);
              s2 = ut(s2, e2), true === s2 && e2.suppressBooleanAttributes ? n2 += ` ${i2.substr(e2.attributeNamePrefix.length)}` : n2 += ` ${i2.substr(e2.attributeNamePrefix.length)}="${s2}"`;
            }
            return n2;
          }
          function lt(t2, e2) {
            let n2 = (t2 = t2.substr(0, t2.length - e2.textNodeName.length - 1)).substr(t2.lastIndexOf(".") + 1);
            for (let i2 in e2.stopNodes) if (e2.stopNodes[i2] === t2 || e2.stopNodes[i2] === "*." + n2) return true;
            return false;
          }
          function ut(t2, e2) {
            if (t2 && t2.length > 0 && e2.processEntities) for (let n2 = 0; n2 < e2.entities.length; n2++) {
              const i2 = e2.entities[n2];
              t2 = t2.replace(i2.regex, i2.val);
            }
            return t2;
          }
          const dt = { attributeNamePrefix: "@_", attributesGroupName: false, textNodeName: "#text", ignoreAttributes: true, cdataPropName: false, format: false, indentBy: "  ", suppressEmptyNode: false, suppressUnpairedNode: true, suppressBooleanAttributes: true, tagValueProcessor: function(t2, e2) {
            return e2;
          }, attributeValueProcessor: function(t2, e2) {
            return e2;
          }, preserveOrder: false, commentPropName: false, unpairedTags: [], entities: [{ regex: new RegExp("&", "g"), val: "&amp;" }, { regex: new RegExp(">", "g"), val: "&gt;" }, { regex: new RegExp("<", "g"), val: "&lt;" }, { regex: new RegExp("'", "g"), val: "&apos;" }, { regex: new RegExp('"', "g"), val: "&quot;" }], processEntities: true, stopNodes: [], oneListGroup: false };
          function ht(t2) {
            var e2;
            this.options = Object.assign({}, dt, t2), true === this.options.ignoreAttributes || this.options.attributesGroupName ? this.isAttribute = function() {
              return false;
            } : (this.ignoreAttributesFn = "function" == typeof (e2 = this.options.ignoreAttributes) ? e2 : Array.isArray(e2) ? (t3) => {
              for (const n2 of e2) {
                if ("string" == typeof n2 && t3 === n2) return true;
                if (n2 instanceof RegExp && n2.test(t3)) return true;
              }
            } : () => false, this.attrPrefixLen = this.options.attributeNamePrefix.length, this.isAttribute = ft), this.processTextOrObjNode = pt, this.options.format ? (this.indentate = ct, this.tagEndChar = ">\n", this.newLine = "\n") : (this.indentate = function() {
              return "";
            }, this.tagEndChar = ">", this.newLine = "");
          }
          function pt(t2, e2, n2, i2) {
            const s2 = this.j2x(t2, n2 + 1, i2.concat(e2));
            return void 0 !== t2[this.options.textNodeName] && 1 === Object.keys(t2).length ? this.buildTextValNode(t2[this.options.textNodeName], e2, s2.attrStr, n2) : this.buildObjectNode(s2.val, e2, s2.attrStr, n2);
          }
          function ct(t2) {
            return this.options.indentBy.repeat(t2);
          }
          function ft(t2) {
            return !(!t2.startsWith(this.options.attributeNamePrefix) || t2 === this.options.textNodeName) && t2.substr(this.attrPrefixLen);
          }
          ht.prototype.build = function(t2) {
            return this.options.preserveOrder ? st(t2, this.options) : (Array.isArray(t2) && this.options.arrayNodeName && this.options.arrayNodeName.length > 1 && (t2 = { [this.options.arrayNodeName]: t2 }), this.j2x(t2, 0, []).val);
          }, ht.prototype.j2x = function(t2, e2, n2) {
            let i2 = "", s2 = "";
            const r2 = n2.join(".");
            for (let o2 in t2) if (Object.prototype.hasOwnProperty.call(t2, o2)) if (void 0 === t2[o2]) this.isAttribute(o2) && (s2 += "");
            else if (null === t2[o2]) this.isAttribute(o2) || o2 === this.options.cdataPropName ? s2 += "" : "?" === o2[0] ? s2 += this.indentate(e2) + "<" + o2 + "?" + this.tagEndChar : s2 += this.indentate(e2) + "<" + o2 + "/" + this.tagEndChar;
            else if (t2[o2] instanceof Date) s2 += this.buildTextValNode(t2[o2], o2, "", e2);
            else if ("object" != typeof t2[o2]) {
              const n3 = this.isAttribute(o2);
              if (n3 && !this.ignoreAttributesFn(n3, r2)) i2 += this.buildAttrPairStr(n3, "" + t2[o2]);
              else if (!n3) if (o2 === this.options.textNodeName) {
                let e3 = this.options.tagValueProcessor(o2, "" + t2[o2]);
                s2 += this.replaceEntitiesValue(e3);
              } else s2 += this.buildTextValNode(t2[o2], o2, "", e2);
            } else if (Array.isArray(t2[o2])) {
              const i3 = t2[o2].length;
              let r3 = "", a2 = "";
              for (let l2 = 0; l2 < i3; l2++) {
                const i4 = t2[o2][l2];
                if (void 0 === i4) ;
                else if (null === i4) "?" === o2[0] ? s2 += this.indentate(e2) + "<" + o2 + "?" + this.tagEndChar : s2 += this.indentate(e2) + "<" + o2 + "/" + this.tagEndChar;
                else if ("object" == typeof i4) if (this.options.oneListGroup) {
                  const t3 = this.j2x(i4, e2 + 1, n2.concat(o2));
                  r3 += t3.val, this.options.attributesGroupName && i4.hasOwnProperty(this.options.attributesGroupName) && (a2 += t3.attrStr);
                } else r3 += this.processTextOrObjNode(i4, o2, e2, n2);
                else if (this.options.oneListGroup) {
                  let t3 = this.options.tagValueProcessor(o2, i4);
                  t3 = this.replaceEntitiesValue(t3), r3 += t3;
                } else r3 += this.buildTextValNode(i4, o2, "", e2);
              }
              this.options.oneListGroup && (r3 = this.buildObjectNode(r3, o2, a2, e2)), s2 += r3;
            } else if (this.options.attributesGroupName && o2 === this.options.attributesGroupName) {
              const e3 = Object.keys(t2[o2]), n3 = e3.length;
              for (let s3 = 0; s3 < n3; s3++) i2 += this.buildAttrPairStr(e3[s3], "" + t2[o2][e3[s3]]);
            } else s2 += this.processTextOrObjNode(t2[o2], o2, e2, n2);
            return { attrStr: i2, val: s2 };
          }, ht.prototype.buildAttrPairStr = function(t2, e2) {
            return e2 = this.options.attributeValueProcessor(t2, "" + e2), e2 = this.replaceEntitiesValue(e2), this.options.suppressBooleanAttributes && "true" === e2 ? " " + t2 : " " + t2 + '="' + e2 + '"';
          }, ht.prototype.buildObjectNode = function(t2, e2, n2, i2) {
            if ("" === t2) return "?" === e2[0] ? this.indentate(i2) + "<" + e2 + n2 + "?" + this.tagEndChar : this.indentate(i2) + "<" + e2 + n2 + this.closeTag(e2) + this.tagEndChar;
            {
              let s2 = "</" + e2 + this.tagEndChar, r2 = "";
              return "?" === e2[0] && (r2 = "?", s2 = ""), !n2 && "" !== n2 || -1 !== t2.indexOf("<") ? false !== this.options.commentPropName && e2 === this.options.commentPropName && 0 === r2.length ? this.indentate(i2) + `<!--${t2}-->` + this.newLine : this.indentate(i2) + "<" + e2 + n2 + r2 + this.tagEndChar + t2 + this.indentate(i2) + s2 : this.indentate(i2) + "<" + e2 + n2 + r2 + ">" + t2 + s2;
            }
          }, ht.prototype.closeTag = function(t2) {
            let e2 = "";
            return -1 !== this.options.unpairedTags.indexOf(t2) ? this.options.suppressUnpairedNode || (e2 = "/") : e2 = this.options.suppressEmptyNode ? "/" : `></${t2}`, e2;
          }, ht.prototype.buildTextValNode = function(t2, e2, n2, i2) {
            if (false !== this.options.cdataPropName && e2 === this.options.cdataPropName) return this.indentate(i2) + `<![CDATA[${t2}]]>` + this.newLine;
            if (false !== this.options.commentPropName && e2 === this.options.commentPropName) return this.indentate(i2) + `<!--${t2}-->` + this.newLine;
            if ("?" === e2[0]) return this.indentate(i2) + "<" + e2 + n2 + "?" + this.tagEndChar;
            {
              let s2 = this.options.tagValueProcessor(e2, t2);
              return s2 = this.replaceEntitiesValue(s2), "" === s2 ? this.indentate(i2) + "<" + e2 + n2 + this.closeTag(e2) + this.tagEndChar : this.indentate(i2) + "<" + e2 + n2 + ">" + s2 + "</" + e2 + this.tagEndChar;
            }
          }, ht.prototype.replaceEntitiesValue = function(t2) {
            if (t2 && t2.length > 0 && this.options.processEntities) for (let e2 = 0; e2 < this.options.entities.length; e2++) {
              const n2 = this.options.entities[e2];
              t2 = t2.replace(n2.regex, n2.val);
            }
            return t2;
          };
          const gt = ht, xt = { validate: a };
          module3.exports = e;
        })();
      }, {}], 2: [function(require2, module3, exports3) {
        var hasOwn = Object.prototype.hasOwnProperty;
        var toString = Object.prototype.toString;
        module3.exports = function forEach(obj, fn, ctx) {
          if (toString.call(fn) !== "[object Function]") {
            throw new TypeError("iterator must be a function");
          }
          var l = obj.length;
          if (l === +l) {
            for (var i = 0; i < l; i++) {
              fn.call(ctx, obj[i], i, obj);
            }
          } else {
            for (var k in obj) {
              if (hasOwn.call(obj, k)) {
                fn.call(ctx, obj[k], k, obj);
              }
            }
          }
        };
      }, {}], 3: [function(require2, module3, exports3) {
        "use strict";
        var each = require2("foreach");
        module3.exports = api;
        function api(obj, pointer, value) {
          if (arguments.length === 3) {
            return api.set(obj, pointer, value);
          }
          if (arguments.length === 2) {
            return api.get(obj, pointer);
          }
          var wrapped = api.bind(api, obj);
          for (var name in api) {
            if (api.hasOwnProperty(name)) {
              wrapped[name] = api[name].bind(wrapped, obj);
            }
          }
          return wrapped;
        }
        api.get = function get(obj, pointer) {
          var refTokens = Array.isArray(pointer) ? pointer : api.parse(pointer);
          for (var i = 0; i < refTokens.length; ++i) {
            var tok = refTokens[i];
            if (!(typeof obj == "object" && tok in obj)) {
              throw new Error("Invalid reference token: " + tok);
            }
            obj = obj[tok];
          }
          return obj;
        };
        api.set = function set(obj, pointer, value) {
          var refTokens = Array.isArray(pointer) ? pointer : api.parse(pointer), nextTok = refTokens[0];
          if (refTokens.length === 0) {
            throw Error("Can not set the root object");
          }
          for (var i = 0; i < refTokens.length - 1; ++i) {
            var tok = refTokens[i];
            if (typeof tok !== "string" && typeof tok !== "number") {
              tok = String(tok);
            }
            if (tok === "__proto__" || tok === "constructor" || tok === "prototype") {
              continue;
            }
            if (tok === "-" && Array.isArray(obj)) {
              tok = obj.length;
            }
            nextTok = refTokens[i + 1];
            if (!(tok in obj)) {
              if (nextTok.match(/^(\d+|-)$/)) {
                obj[tok] = [];
              } else {
                obj[tok] = {};
              }
            }
            obj = obj[tok];
          }
          if (nextTok === "-" && Array.isArray(obj)) {
            nextTok = obj.length;
          }
          obj[nextTok] = value;
          return this;
        };
        api.remove = function(obj, pointer) {
          var refTokens = Array.isArray(pointer) ? pointer : api.parse(pointer);
          var finalToken = refTokens[refTokens.length - 1];
          if (finalToken === void 0) {
            throw new Error('Invalid JSON pointer for remove: "' + pointer + '"');
          }
          var parent = api.get(obj, refTokens.slice(0, -1));
          if (Array.isArray(parent)) {
            var index = +finalToken;
            if (finalToken === "" && isNaN(index)) {
              throw new Error('Invalid array index: "' + finalToken + '"');
            }
            Array.prototype.splice.call(parent, index, 1);
          } else {
            delete parent[finalToken];
          }
        };
        api.dict = function dict(obj, descend) {
          var results = {};
          api.walk(obj, function(value, pointer) {
            results[pointer] = value;
          }, descend);
          return results;
        };
        api.walk = function walk(obj, iterator, descend) {
          var refTokens = [];
          descend = descend || function(value) {
            var type = Object.prototype.toString.call(value);
            return type === "[object Object]" || type === "[object Array]";
          };
          (function next(cur) {
            each(cur, function(value, key) {
              refTokens.push(String(key));
              if (descend(value)) {
                next(value);
              } else {
                iterator(value, api.compile(refTokens));
              }
              refTokens.pop();
            });
          })(obj);
        };
        api.has = function has(obj, pointer) {
          try {
            api.get(obj, pointer);
          } catch (e) {
            return false;
          }
          return true;
        };
        api.escape = function escape(str) {
          return str.toString().replace(/~/g, "~0").replace(/\//g, "~1");
        };
        api.unescape = function unescape(str) {
          return str.replace(/~1/g, "/").replace(/~0/g, "~");
        };
        api.parse = function parse(pointer) {
          if (pointer === "") {
            return [];
          }
          if (pointer.charAt(0) !== "/") {
            throw new Error("Invalid JSON pointer: " + pointer);
          }
          return pointer.substring(1).split(/\//).map(api.unescape);
        };
        api.compile = function compile(refTokens) {
          if (refTokens.length === 0) {
            return "";
          }
          return "/" + refTokens.map(api.escape).join("/");
        };
      }, { "foreach": 2 }], 4: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.allOfSample = allOfSample;
        var _traverse = require2("./traverse");
        var _utils = require2("./utils");
        function ownKeys(e, r) {
          var t = Object.keys(e);
          if (Object.getOwnPropertySymbols) {
            var o = Object.getOwnPropertySymbols(e);
            r && (o = o.filter(function(r2) {
              return Object.getOwnPropertyDescriptor(e, r2).enumerable;
            })), t.push.apply(t, o);
          }
          return t;
        }
        function _objectSpread(e) {
          for (var r = 1; r < arguments.length; r++) {
            var t = null != arguments[r] ? arguments[r] : {};
            r % 2 ? ownKeys(Object(t), true).forEach(function(r2) {
              _defineProperty(e, r2, t[r2]);
            }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : ownKeys(Object(t)).forEach(function(r2) {
              Object.defineProperty(e, r2, Object.getOwnPropertyDescriptor(t, r2));
            });
          }
          return e;
        }
        function _defineProperty(e, r, t) {
          return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: true, configurable: true, writable: true }) : e[r] = t, e;
        }
        function _toPropertyKey(t) {
          var i = _toPrimitive(t, "string");
          return "symbol" == typeof i ? i : i + "";
        }
        function _toPrimitive(t, r) {
          if ("object" != typeof t || !t) return t;
          var e = t[Symbol.toPrimitive];
          if (void 0 !== e) {
            var i = e.call(t, r || "default");
            if ("object" != typeof i) return i;
            throw new TypeError("@@toPrimitive must return a primitive value.");
          }
          return ("string" === r ? String : Number)(t);
        }
        function allOfSample(into, children, options, spec, context) {
          let res = (0, _traverse.traverse)(into, options, spec);
          const subSamples = [];
          for (let subSchema of children) {
            const {
              type,
              readOnly,
              writeOnly,
              value
            } = (0, _traverse.traverse)(_objectSpread({
              type: res.type
            }, subSchema), options, spec, _objectSpread(_objectSpread({}, context), {}, {
              isAllOfChild: true
            }));
            if (res.type && type && type !== res.type) {
              console.warn("allOf: schemas with different types can't be merged");
              res.type = type;
            }
            res.type = res.type || type;
            res.readOnly = res.readOnly || readOnly;
            res.writeOnly = res.writeOnly || writeOnly;
            if (value != null) subSamples.push(value);
          }
          if (res.type === "object") {
            res.value = (0, _utils.mergeDeep)(res.value || {}, ...subSamples.filter((sample) => typeof sample === "object"));
            for (const key in res.value) {
              if (res.value[key] === _utils.SKIP_SYMBOL) {
                delete res.value[key];
              }
            }
            return res;
          } else {
            if (res.type === "array") {
              if (!options.quiet) console.warn('OpenAPI Sampler: found allOf with "array" type. Result may be incorrect');
            }
            const lastSample = subSamples[subSamples.length - 1];
            res.value = lastSample != null ? lastSample : res.value;
            return res;
          }
        }
      }, { "./traverse": 14, "./utils": 15 }], 5: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.inferType = inferType;
        const schemaKeywordTypes = {
          multipleOf: "number",
          maximum: "number",
          exclusiveMaximum: "number",
          minimum: "number",
          exclusiveMinimum: "number",
          maxLength: "string",
          minLength: "string",
          pattern: "string",
          items: "array",
          maxItems: "array",
          minItems: "array",
          uniqueItems: "array",
          additionalItems: "array",
          maxProperties: "object",
          minProperties: "object",
          required: "object",
          additionalProperties: "object",
          properties: "object",
          patternProperties: "object",
          dependencies: "object"
        };
        function inferType(schema) {
          if (schema.type !== void 0) {
            return Array.isArray(schema.type) ? schema.type.length === 0 ? null : schema.type[0] : schema.type;
          }
          const keywords = Object.keys(schemaKeywordTypes);
          for (var i = 0; i < keywords.length; i++) {
            let keyword = keywords[i];
            let type = schemaKeywordTypes[keyword];
            if (schema[keyword] !== void 0) {
              return type;
            }
          }
          return null;
        }
      }, {}], 6: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3._registerSampler = _registerSampler;
        exports3._samplers = void 0;
        Object.defineProperty(exports3, "inferType", {
          enumerable: true,
          get: function() {
            return _infer.inferType;
          }
        });
        exports3.sample = sample;
        var _traverse = require2("./traverse");
        var _index = require2("./samplers/index");
        var _fastXmlParser = require2("fast-xml-parser");
        var _infer = require2("./infer");
        var _samplers = exports3._samplers = {};
        const defaults = {
          skipReadOnly: false,
          maxSampleDepth: 15
        };
        function convertJsonToXml(obj, schema) {
          if (!obj) {
            throw new Error("Unknown format output for building XML.");
          }
          if (Array.isArray(obj) || Object.keys(obj).length > 1) {
            var _schema$xml;
            obj = {
              [(schema === null || schema === void 0 || (_schema$xml = schema.xml) === null || _schema$xml === void 0 ? void 0 : _schema$xml.name) || "root"]: obj
            };
          }
          const builder = new _fastXmlParser.XMLBuilder({
            ignoreAttributes: false,
            format: true,
            attributeNamePrefix: "$",
            textNodeName: "#text",
            cdataPropName: "#cdata"
          });
          return builder.build(obj);
        }
        function sample(schema, options, spec) {
          let opts = Object.assign({}, defaults, options);
          (0, _traverse.clearCache)();
          let result = (0, _traverse.traverse)(schema, opts, spec).value;
          if ((opts === null || opts === void 0 ? void 0 : opts.format) === "xml") {
            return convertJsonToXml(result, schema);
          }
          return result;
        }
        ;
        function _registerSampler(type, sampler) {
          _samplers[type] = sampler;
        }
        ;
        _registerSampler("array", _index.sampleArray);
        _registerSampler("boolean", _index.sampleBoolean);
        _registerSampler("integer", _index.sampleNumber);
        _registerSampler("number", _index.sampleNumber);
        _registerSampler("object", _index.sampleObject);
        _registerSampler("string", _index.sampleString);
      }, { "./infer": 5, "./samplers/index": 9, "./traverse": 14, "fast-xml-parser": 1 }], 7: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.sampleArray = sampleArray;
        var _traverse = require2("../traverse");
        var _utils = require2("../utils");
        function ownKeys(e, r) {
          var t = Object.keys(e);
          if (Object.getOwnPropertySymbols) {
            var o = Object.getOwnPropertySymbols(e);
            r && (o = o.filter(function(r2) {
              return Object.getOwnPropertyDescriptor(e, r2).enumerable;
            })), t.push.apply(t, o);
          }
          return t;
        }
        function _objectSpread(e) {
          for (var r = 1; r < arguments.length; r++) {
            var t = null != arguments[r] ? arguments[r] : {};
            r % 2 ? ownKeys(Object(t), true).forEach(function(r2) {
              _defineProperty(e, r2, t[r2]);
            }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : ownKeys(Object(t)).forEach(function(r2) {
              Object.defineProperty(e, r2, Object.getOwnPropertyDescriptor(t, r2));
            });
          }
          return e;
        }
        function _defineProperty(e, r, t) {
          return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: true, configurable: true, writable: true }) : e[r] = t, e;
        }
        function _toPropertyKey(t) {
          var i = _toPrimitive(t, "string");
          return "symbol" == typeof i ? i : i + "";
        }
        function _toPrimitive(t, r) {
          if ("object" != typeof t || !t) return t;
          var e = t[Symbol.toPrimitive];
          if (void 0 !== e) {
            var i = e.call(t, r || "default");
            if ("object" != typeof i) return i;
            throw new TypeError("@@toPrimitive must return a primitive value.");
          }
          return ("string" === r ? String : Number)(t);
        }
        function sampleArray(schema) {
          let options = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};
          let spec = arguments.length > 2 ? arguments[2] : void 0;
          let context = arguments.length > 3 ? arguments[3] : void 0;
          const depth = context && context.depth || 1;
          let arrayLength = Math.min(schema.maxItems != void 0 ? schema.maxItems : Infinity, schema.minItems || 1);
          const items = schema.prefixItems || schema.items || schema.contains;
          if (Array.isArray(items)) {
            arrayLength = Math.max(arrayLength, items.length);
          }
          let itemSchemaGetter = (itemNumber) => {
            if (Array.isArray(items)) {
              return items[itemNumber] || {};
            }
            return items || {};
          };
          let res = [];
          if (!items) return res;
          for (let i = 0; i < arrayLength; i++) {
            let itemSchema = itemSchemaGetter(i);
            let {
              value: sample
            } = (0, _traverse.traverse)(itemSchema, options, spec, {
              depth: depth + 1
            });
            if ((options === null || options === void 0 ? void 0 : options.format) === "xml") {
              const {
                value,
                propertyName
              } = (0, _utils.applyXMLAttributes)({
                value: sample
              }, itemSchema, context);
              if (propertyName) {
                var _res;
                if (!((_res = res) !== null && _res !== void 0 && _res[propertyName])) {
                  res = _objectSpread(_objectSpread({}, res), {}, {
                    [propertyName]: []
                  });
                }
                res[propertyName].push(value);
              } else {
                res = _objectSpread(_objectSpread({}, res), value);
              }
            } else {
              res.push(sample);
            }
          }
          if ((options === null || options === void 0 ? void 0 : options.format) === "xml" && depth === 1) {
            const {
              value,
              propertyName
            } = (0, _utils.applyXMLAttributes)({
              value: null
            }, schema, context);
            if (propertyName) {
              if (value) {
                res = Array.isArray(res) ? {
                  [propertyName]: _objectSpread(_objectSpread({}, value), res.map((item) => ({
                    ["#text"]: _objectSpread({}, item)
                  })))
                } : {
                  [propertyName]: _objectSpread(_objectSpread({}, res), value)
                };
              } else {
                res = {
                  [propertyName]: res
                };
              }
            }
          }
          return res;
        }
      }, { "../traverse": 14, "../utils": 15 }], 8: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.sampleBoolean = sampleBoolean;
        function sampleBoolean(schema) {
          return true;
        }
      }, {}], 9: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        Object.defineProperty(exports3, "sampleArray", {
          enumerable: true,
          get: function() {
            return _array.sampleArray;
          }
        });
        Object.defineProperty(exports3, "sampleBoolean", {
          enumerable: true,
          get: function() {
            return _boolean.sampleBoolean;
          }
        });
        Object.defineProperty(exports3, "sampleNumber", {
          enumerable: true,
          get: function() {
            return _number.sampleNumber;
          }
        });
        Object.defineProperty(exports3, "sampleObject", {
          enumerable: true,
          get: function() {
            return _object.sampleObject;
          }
        });
        Object.defineProperty(exports3, "sampleString", {
          enumerable: true,
          get: function() {
            return _string.sampleString;
          }
        });
        var _array = require2("./array");
        var _boolean = require2("./boolean");
        var _number = require2("./number");
        var _object = require2("./object");
        var _string = require2("./string");
      }, { "./array": 7, "./boolean": 8, "./number": 10, "./object": 11, "./string": 13 }], 10: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.sampleNumber = sampleNumber;
        function sampleNumber(schema) {
          let res = 0;
          if (schema.type === "number" && (schema.format === "float" || schema.format === "double")) {
            res = 0.1;
          }
          if (typeof schema.exclusiveMinimum === "boolean" || typeof schema.exclusiveMaximum === "boolean") {
            if (schema.maximum && schema.minimum) {
              res = schema.exclusiveMinimum ? Math.floor(schema.minimum) + 1 : schema.minimum;
              if (schema.exclusiveMaximum && res >= schema.maximum || !schema.exclusiveMaximum && res > schema.maximum) {
                res = (schema.maximum + schema.minimum) / 2;
              }
              return res;
            }
            if (schema.minimum) {
              if (schema.exclusiveMinimum) {
                return Math.floor(schema.minimum) + 1;
              } else {
                return schema.minimum;
              }
            }
            if (schema.maximum) {
              if (schema.exclusiveMaximum) {
                return schema.maximum > 0 ? 0 : Math.floor(schema.maximum) - 1;
              } else {
                return schema.maximum > 0 ? 0 : schema.maximum;
              }
            }
          } else {
            if (schema.minimum) {
              return schema.minimum;
            }
            if (schema.exclusiveMinimum) {
              res = Math.floor(schema.exclusiveMinimum) + 1;
              if (res === schema.exclusiveMaximum) {
                res = (res + Math.floor(schema.exclusiveMaximum) - 1) / 2;
              }
            } else if (schema.exclusiveMaximum) {
              res = Math.floor(schema.exclusiveMaximum) - 1;
            } else if (schema.maximum) {
              res = schema.maximum;
            }
          }
          return res;
        }
      }, {}], 11: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.sampleObject = sampleObject;
        var _traverse = require2("../traverse");
        var _utils = require2("../utils");
        function ownKeys(e, r) {
          var t = Object.keys(e);
          if (Object.getOwnPropertySymbols) {
            var o = Object.getOwnPropertySymbols(e);
            r && (o = o.filter(function(r2) {
              return Object.getOwnPropertyDescriptor(e, r2).enumerable;
            })), t.push.apply(t, o);
          }
          return t;
        }
        function _objectSpread(e) {
          for (var r = 1; r < arguments.length; r++) {
            var t = null != arguments[r] ? arguments[r] : {};
            r % 2 ? ownKeys(Object(t), true).forEach(function(r2) {
              _defineProperty(e, r2, t[r2]);
            }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : ownKeys(Object(t)).forEach(function(r2) {
              Object.defineProperty(e, r2, Object.getOwnPropertyDescriptor(t, r2));
            });
          }
          return e;
        }
        function _defineProperty(e, r, t) {
          return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: true, configurable: true, writable: true }) : e[r] = t, e;
        }
        function _toPropertyKey(t) {
          var i = _toPrimitive(t, "string");
          return "symbol" == typeof i ? i : i + "";
        }
        function _toPrimitive(t, r) {
          if ("object" != typeof t || !t) return t;
          var e = t[Symbol.toPrimitive];
          if (void 0 !== e) {
            var i = e.call(t, r || "default");
            if ("object" != typeof i) return i;
            throw new TypeError("@@toPrimitive must return a primitive value.");
          }
          return ("string" === r ? String : Number)(t);
        }
        function sampleObject(schema) {
          let options = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};
          let spec = arguments.length > 2 ? arguments[2] : void 0;
          let context = arguments.length > 3 ? arguments[3] : void 0;
          let res = {};
          const depth = context && context.depth || 1;
          if (schema && typeof schema.properties === "object") {
            const requiredProperties = Array.isArray(schema.required) ? schema.required : [];
            const requiredPropertiesMap = {};
            for (const requiredProperty of requiredProperties) {
              requiredPropertiesMap[requiredProperty] = true;
            }
            Object.keys(schema.properties).forEach((propertyName) => {
              if (options.skipNonRequired && !requiredPropertiesMap.hasOwnProperty(propertyName)) {
                return;
              }
              const sample = (0, _traverse.traverse)(schema.properties[propertyName], options, spec, {
                propertyName,
                depth: depth + 1
              });
              if (options.skipReadOnly && sample.readOnly) {
                if (context !== null && context !== void 0 && context.isAllOfChild) {
                  res[propertyName] = _utils.SKIP_SYMBOL;
                }
                return;
              }
              if (options.skipWriteOnly && sample.writeOnly) {
                if (context !== null && context !== void 0 && context.isAllOfChild) {
                  res[propertyName] = _utils.SKIP_SYMBOL;
                }
                return;
              }
              if ((options === null || options === void 0 ? void 0 : options.format) === "xml") {
                const {
                  propertyName: newPropertyName,
                  value
                } = (0, _utils.applyXMLAttributes)(sample, schema.properties[propertyName], {
                  propertyName
                });
                if (newPropertyName) {
                  res[newPropertyName] = value;
                } else if (value !== null && typeof value === "object") {
                  res = _objectSpread(_objectSpread({}, res), value);
                }
              } else {
                res[propertyName] = sample.value;
              }
            });
          }
          if (schema && typeof schema.additionalProperties === "object") {
            const propertyName = schema.additionalProperties["x-additionalPropertiesName"] || "property";
            res["".concat(String(propertyName), "1")] = (0, _traverse.traverse)(schema.additionalProperties, options, spec, {
              depth: depth + 1
            }).value;
            res["".concat(String(propertyName), "2")] = (0, _traverse.traverse)(schema.additionalProperties, options, spec, {
              depth: depth + 1
            }).value;
          }
          if (schema && typeof schema.properties === "object" && schema.maxProperties !== void 0 && Object.keys(res).length > schema.maxProperties) {
            const filteredResult = {};
            let propertiesAdded = 0;
            const requiredProperties = Array.isArray(schema.required) ? schema.required : [];
            requiredProperties.forEach((propName) => {
              if (res[propName] !== void 0) {
                filteredResult[propName] = res[propName];
                propertiesAdded++;
              }
            });
            Object.keys(res).forEach((propName) => {
              if (propertiesAdded < schema.maxProperties && !filteredResult.hasOwnProperty(propName)) {
                filteredResult[propName] = res[propName];
                propertiesAdded++;
              }
            });
            res = filteredResult;
          }
          return res;
        }
      }, { "../traverse": 14, "../utils": 15 }], 12: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.regexSample = regexSample;
        function boolSample() {
          return true;
        }
        function intSample(min, _max) {
          return min;
        }
        function getRepetitionsBasedOnQuantifierParameters(quantifierSymbol, quantifierMin, quantifierMax) {
          let repetitions = 1;
          if (quantifierSymbol) {
            switch (quantifierSymbol) {
              case "?": {
                repetitions = boolSample() ? 0 : 1;
                break;
              }
              case "*": {
                const limit = 8;
                repetitions = intSample(0, limit);
                break;
              }
              case "+": {
                const limit = 8;
                repetitions = intSample(1, limit);
                break;
              }
              default:
                throw new Error("Unknown quantifier symbol provided.");
            }
          } else if (quantifierMin != null && quantifierMax != null) {
            repetitions = intSample(parseInt(quantifierMin), parseInt(quantifierMax));
          } else if (quantifierMin != null && quantifierMax == null) {
            repetitions = parseInt(quantifierMin);
          }
          return repetitions;
        }
        function regexSample(pattern) {
          let isCaseInsensitive = false;
          if (pattern instanceof RegExp) {
            var _pattern$match$, _pattern$match;
            isCaseInsensitive = pattern.flags.includes("i");
            pattern = pattern.toString();
            pattern = (_pattern$match$ = (_pattern$match = pattern.match(/\/(.+?)\//)) === null || _pattern$match === void 0 ? void 0 : _pattern$match[1]) !== null && _pattern$match$ !== void 0 ? _pattern$match$ : "";
          }
          pattern = pattern.replace(/^(\^)?(.*?)(\$)?$/, "$2");
          let min;
          let max;
          let repetitions;
          const SINGLE_CHAR_REG = /([.A-Za-z0-9])(?:\{(\d+)(?:\,(\d+)|)\}|(\?|\*|\+))(?![^[]*]|[^{]*})/;
          let token = pattern.match(SINGLE_CHAR_REG);
          while (token != null) {
            const quantifierMin = token[2];
            const quantifierMax = token[3];
            const quantifierSymbol = token[4];
            repetitions = getRepetitionsBasedOnQuantifierParameters(quantifierSymbol, quantifierMin, quantifierMax);
            pattern = pattern.slice(0, token.index) + token[1].repeat(repetitions) + pattern.slice(token.index + token[0].length);
            token = pattern.match(SINGLE_CHAR_REG);
          }
          const SINGLE_RANGE_REG = /(\d-\d|\w-\w|\d|\w|[-!@#$&()`.+,/"])/;
          const RANGE_ALPHANUMEMRIC_REG = /\[(\^|)(-|)(.+?)\](?:\{(\d+)(?:\,(\d+)|)\}|(\?|\*|\+)|)/;
          token = pattern.match(RANGE_ALPHANUMEMRIC_REG);
          while (token != null) {
            const isNegated = token[1] === "^";
            const includesDash = token[2] === "-";
            const quantifierMin = token[4];
            const quantifierMax = token[5];
            const quantifierSymbol = token[6];
            const rangeCodes = [];
            let ranges = token[3];
            let range = ranges.match(SINGLE_RANGE_REG);
            if (includesDash) {
              rangeCodes.push(45);
            }
            while (range != null) {
              if (range[0].indexOf("-") === -1) {
                if (isCaseInsensitive && isNaN(Number(range[0]))) {
                  rangeCodes.push(range[0].toUpperCase().charCodeAt(0));
                  rangeCodes.push(range[0].toLowerCase().charCodeAt(0));
                } else {
                  rangeCodes.push(range[0].charCodeAt(0));
                }
              } else {
                const rangeMinMax = range[0].split("-").map((x) => x.charCodeAt(0));
                min = rangeMinMax[0];
                max = rangeMinMax[1];
                if (min > max) {
                  throw new Error("Character range provided is out of order.");
                }
                for (let i = min; i <= max; i++) {
                  if (isCaseInsensitive && isNaN(Number(String.fromCharCode(i)))) {
                    const ch = String.fromCharCode(i);
                    rangeCodes.push(ch.toUpperCase().charCodeAt(0));
                    rangeCodes.push(ch.toLowerCase().charCodeAt(0));
                  } else {
                    rangeCodes.push(i);
                  }
                }
              }
              ranges = ranges.substring(range[0].length);
              range = ranges.match(SINGLE_RANGE_REG);
            }
            repetitions = getRepetitionsBasedOnQuantifierParameters(quantifierSymbol, quantifierMin, quantifierMax);
            if (isNegated) {
              let index = -1;
              for (let i = 48; i <= 57; i++) {
                index = rangeCodes.indexOf(i);
                if (index > -1) {
                  rangeCodes.splice(index, 1);
                  continue;
                }
                rangeCodes.push(i);
              }
              for (let i = 65; i <= 90; i++) {
                index = rangeCodes.indexOf(i);
                if (index > -1) {
                  rangeCodes.splice(index, 1);
                  continue;
                }
                rangeCodes.push(i);
              }
              for (let i = 97; i <= 122; i++) {
                index = rangeCodes.indexOf(i);
                if (index > -1) {
                  rangeCodes.splice(index, 1);
                  continue;
                }
                rangeCodes.push(i);
              }
            }
            const generatedString = Array.from({
              length: repetitions
            }, () => String.fromCharCode(rangeCodes[intSample(0, rangeCodes.length - 1)])).join("");
            pattern = pattern.slice(0, token.index) + generatedString + pattern.slice(token.index + token[0].length);
            token = pattern.match(RANGE_ALPHANUMEMRIC_REG);
          }
          const RANGE_REP_REG = /(.)\{(\d+)\,(\d+)\}/;
          token = pattern.match(RANGE_REP_REG);
          while (token != null) {
            min = parseInt(token[2]);
            max = parseInt(token[3]);
            if (min > max) {
              throw new Error("Numbers out of order in {} quantifier.");
            }
            repetitions = intSample(min, max);
            pattern = pattern.slice(0, token.index) + token[1].repeat(repetitions) + pattern.slice(token.index + token[0].length);
            token = pattern.match(RANGE_REP_REG);
          }
          const REP_REG = /(.)\{(\d+)\}/;
          token = pattern.match(REP_REG);
          while (token != null) {
            repetitions = parseInt(token[2]);
            pattern = pattern.slice(0, token.index) + token[1].repeat(repetitions) + pattern.slice(token.index + token[0].length);
            token = pattern.match(REP_REG);
          }
          return pattern;
        }
      }, {}], 13: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.sampleString = sampleString;
        var _utils = require2("../utils");
        var faker = _interopRequireWildcard(require2("./string-regex"));
        function _interopRequireWildcard(e, t) {
          if ("function" == typeof WeakMap) var r = /* @__PURE__ */ new WeakMap(), n = /* @__PURE__ */ new WeakMap();
          return (_interopRequireWildcard = function(e2, t2) {
            if (!t2 && e2 && e2.__esModule) return e2;
            var o, i, f = { __proto__: null, default: e2 };
            if (null === e2 || "object" != typeof e2 && "function" != typeof e2) return f;
            if (o = t2 ? n : r) {
              if (o.has(e2)) return o.get(e2);
              o.set(e2, f);
            }
            for (const t3 in e2) "default" !== t3 && {}.hasOwnProperty.call(e2, t3) && ((i = (o = Object.defineProperty) && Object.getOwnPropertyDescriptor(e2, t3)) && (i.get || i.set) ? o(f, t3, i) : f[t3] = e2[t3]);
            return f;
          })(e, t);
        }
        const passwordSymbols = "qwerty!@#$%^123456";
        function emailSample() {
          return "user@example.com";
        }
        function idnEmailSample() {
          return "\u043F\u043E\u0448\u0442\u0430@\u0443\u043A\u0440.\u043D\u0435\u0442";
        }
        function passwordSample(min, max) {
          let res = "pa$$word";
          if (min > res.length) {
            res += "_";
            res += (0, _utils.ensureMinLength)(passwordSymbols, min - res.length).substring(0, min - res.length);
          }
          return res;
        }
        function commonDateTimeSample(_ref) {
          let {
            min,
            max,
            omitTime,
            omitDate
          } = _ref;
          let res = (0, _utils.toRFCDateTime)(/* @__PURE__ */ new Date("2019-08-24T14:15:22.123Z"), omitTime, omitDate, false);
          if (res.length < min) {
            console.warn("Using minLength = ".concat(min, ' is incorrect with format "date-time"'));
          }
          if (max && res.length > max) {
            console.warn("Using maxLength = ".concat(max, ' is incorrect with format "date-time"'));
          }
          return res;
        }
        function dateTimeSample(min, max) {
          return commonDateTimeSample({
            min,
            max,
            omitTime: false,
            omitDate: false
          });
        }
        function dateSample(min, max) {
          return commonDateTimeSample({
            min,
            max,
            omitTime: true,
            omitDate: false
          });
        }
        function timeSample(min, max) {
          return commonDateTimeSample({
            min,
            max,
            omitTime: false,
            omitDate: true
          }).slice(1);
        }
        function defaultSample(min, max, _propertyName, pattern) {
          let enablePatterns = arguments.length > 4 && arguments[4] !== void 0 ? arguments[4] : false;
          if (pattern && enablePatterns) {
            return faker.regexSample(pattern);
          }
          let res = (0, _utils.ensureMinLength)("string", min);
          if (max && res.length > max) {
            res = res.substring(0, max);
          }
          return res;
        }
        function ipv4Sample() {
          return "192.168.0.1";
        }
        function ipv6Sample() {
          return "2001:0db8:85a3:0000:0000:8a2e:0370:7334";
        }
        function hostnameSample() {
          return "example.com";
        }
        function idnHostnameSample() {
          return "\u043F\u0440\u0438\u043A\u043B\u0430\u0434.\u0443\u043A\u0440";
        }
        function uriSample() {
          return "http://example.com";
        }
        function uriReferenceSample() {
          return "../dictionary";
        }
        function uriTemplateSample() {
          return "http://example.com/{endpoint}";
        }
        function iriSample() {
          return "http://example.com/entity/1";
        }
        function iriReferenceSample() {
          return "/entity/1";
        }
        function uuidSample(_min, _max, propertyName) {
          return (0, _utils.uuid)(propertyName || "id");
        }
        function jsonPointerSample() {
          return "/json/pointer";
        }
        function relativeJsonPointerSample() {
          return "1/relative/json/pointer";
        }
        function regexSample() {
          return "/regex/";
        }
        const stringFormats = {
          "email": emailSample,
          "idn-email": idnEmailSample,
          // https://tools.ietf.org/html/rfc6531#section-3.3
          "password": passwordSample,
          "date-time": dateTimeSample,
          "date": dateSample,
          "time": timeSample,
          // full-time in https://tools.ietf.org/html/rfc3339#section-5.6
          "ipv4": ipv4Sample,
          "ipv6": ipv6Sample,
          "hostname": hostnameSample,
          "idn-hostname": idnHostnameSample,
          // https://tools.ietf.org/html/rfc5890#section-2.3.2.3
          "iri": iriSample,
          // https://tools.ietf.org/html/rfc3987
          "iri-reference": iriReferenceSample,
          "uri": uriSample,
          "uri-reference": uriReferenceSample,
          // either a URI or relative-reference https://tools.ietf.org/html/rfc3986#section-4.1
          "uri-template": uriTemplateSample,
          "uuid": uuidSample,
          "default": defaultSample,
          "json-pointer": jsonPointerSample,
          "relative-json-pointer": relativeJsonPointerSample,
          // https://tools.ietf.org/html/draft-handrews-relative-json-pointer-01
          "regex": regexSample
        };
        function sampleString(schema, options, spec, context) {
          let format = schema.format || "default";
          let sampler = stringFormats[format] || defaultSample;
          let propertyName = context && context.propertyName;
          return sampler(schema.minLength || 0, schema.maxLength, propertyName, schema.pattern, options === null || options === void 0 ? void 0 : options.enablePatterns);
        }
      }, { "../utils": 15, "./string-regex": 12 }], 14: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.clearCache = clearCache;
        exports3.traverse = traverse;
        var _openapiSampler = require2("./openapi-sampler");
        var _allOf = require2("./allOf");
        var _infer = require2("./infer");
        var _utils = require2("./utils");
        var _jsonPointer = _interopRequireDefault(require2("json-pointer"));
        const _excluded = ["if", "then"];
        function _interopRequireDefault(e) {
          return e && e.__esModule ? e : { default: e };
        }
        function _objectWithoutProperties(e, t) {
          if (null == e) return {};
          var o, r, i = _objectWithoutPropertiesLoose(e, t);
          if (Object.getOwnPropertySymbols) {
            var n = Object.getOwnPropertySymbols(e);
            for (r = 0; r < n.length; r++) o = n[r], -1 === t.indexOf(o) && {}.propertyIsEnumerable.call(e, o) && (i[o] = e[o]);
          }
          return i;
        }
        function _objectWithoutPropertiesLoose(r, e) {
          if (null == r) return {};
          var t = {};
          for (var n in r) if ({}.hasOwnProperty.call(r, n)) {
            if (-1 !== e.indexOf(n)) continue;
            t[n] = r[n];
          }
          return t;
        }
        function ownKeys(e, r) {
          var t = Object.keys(e);
          if (Object.getOwnPropertySymbols) {
            var o = Object.getOwnPropertySymbols(e);
            r && (o = o.filter(function(r2) {
              return Object.getOwnPropertyDescriptor(e, r2).enumerable;
            })), t.push.apply(t, o);
          }
          return t;
        }
        function _objectSpread(e) {
          for (var r = 1; r < arguments.length; r++) {
            var t = null != arguments[r] ? arguments[r] : {};
            r % 2 ? ownKeys(Object(t), true).forEach(function(r2) {
              _defineProperty(e, r2, t[r2]);
            }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : ownKeys(Object(t)).forEach(function(r2) {
              Object.defineProperty(e, r2, Object.getOwnPropertyDescriptor(t, r2));
            });
          }
          return e;
        }
        function _defineProperty(e, r, t) {
          return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: true, configurable: true, writable: true }) : e[r] = t, e;
        }
        function _toPropertyKey(t) {
          var i = _toPrimitive(t, "string");
          return "symbol" == typeof i ? i : i + "";
        }
        function _toPrimitive(t, r) {
          if ("object" != typeof t || !t) return t;
          var e = t[Symbol.toPrimitive];
          if (void 0 !== e) {
            var i = e.call(t, r || "default");
            if ("object" != typeof i) return i;
            throw new TypeError("@@toPrimitive must return a primitive value.");
          }
          return ("string" === r ? String : Number)(t);
        }
        let $refCache = {};
        let seenSchemasStack = [];
        function clearCache() {
          $refCache = {};
          seenSchemasStack = [];
        }
        function inferExample(schema) {
          let example;
          if (schema.const !== void 0) {
            example = schema.const;
          } else if (schema.examples !== void 0 && schema.examples.length) {
            example = schema.examples[0];
          } else if (schema.enum !== void 0 && schema.enum.length) {
            example = schema.enum[0];
          } else if (schema.default !== void 0) {
            example = schema.default;
          }
          return example;
        }
        function tryInferExample(schema) {
          const example = inferExample(schema);
          if (example !== void 0) {
            return {
              value: example,
              readOnly: schema.readOnly,
              writeOnly: schema.writeOnly,
              type: null
            };
          }
          return;
        }
        function traverse(schema, options, spec, context) {
          if (context) {
            if (seenSchemasStack.includes(schema)) return (0, _utils.getResultForCircular)((0, _infer.inferType)(schema));
            seenSchemasStack.push(schema);
          }
          if (context && context.depth > options.maxSampleDepth) {
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            return (0, _utils.getResultForCircular)((0, _infer.inferType)(schema));
          }
          if (schema.$ref) {
            if (!spec) {
              throw new Error("Your schema contains $ref. You must provide full specification in the third parameter.");
            }
            let ref = decodeURIComponent(schema.$ref);
            if (ref.startsWith("#")) {
              ref = ref.substring(1);
            }
            const referenced = _jsonPointer.default.get(spec, ref);
            let result;
            if ($refCache[ref] !== true) {
              $refCache[ref] = true;
              const traverseResult = traverse(referenced, options, spec, context);
              if (options.format === "xml") {
                const refName = ref.split("/").pop();
                const xmlContext = _objectSpread(_objectSpread({}, context), {}, {
                  propertyName: (context === null || context === void 0 ? void 0 : context.propertyName) || refName
                });
                const {
                  propertyName,
                  value
                } = (0, _utils.applyXMLAttributes)(traverseResult, referenced, xmlContext);
                result = _objectSpread(_objectSpread({}, traverseResult), {}, {
                  value: {
                    [propertyName || "root"]: value
                  }
                });
              } else {
                result = traverseResult;
              }
              $refCache[ref] = false;
            } else {
              const referencedType = (0, _infer.inferType)(referenced);
              result = (0, _utils.getResultForCircular)(referencedType);
            }
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            return result;
          }
          if (schema.example !== void 0) {
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            return {
              value: schema.example,
              readOnly: schema.readOnly,
              writeOnly: schema.writeOnly,
              type: schema.type
            };
          }
          if (schema.allOf !== void 0) {
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            return tryInferExample(schema) || (0, _allOf.allOfSample)(_objectSpread(_objectSpread({}, schema), {}, {
              allOf: void 0
            }), schema.allOf, options, spec, context);
          }
          if (schema.oneOf && schema.oneOf.length) {
            if (schema.anyOf) {
              if (!options.quiet) console.warn("oneOf and anyOf are not supported on the same level. Skipping anyOf");
            }
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            const firstOneOf = Object.assign({
              readOnly: schema.readOnly,
              writeOnly: schema.writeOnly
            }, schema.oneOf[0]);
            return traverseOneOrAnyOf(schema, firstOneOf);
          }
          if (schema.anyOf && schema.anyOf.length) {
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            const firstAnyOf = Object.assign({
              readOnly: schema.readOnly,
              writeOnly: schema.writeOnly
            }, schema.anyOf[0]);
            return traverseOneOrAnyOf(schema, firstAnyOf);
          }
          if (schema.if && schema.then) {
            (0, _utils.popSchemaStack)(seenSchemasStack, context);
            const {
              if: ifSchema,
              then
            } = schema, rest = _objectWithoutProperties(schema, _excluded);
            return traverse((0, _utils.mergeDeep)(rest, ifSchema, then), options, spec, context);
          }
          let example = inferExample(schema);
          let type = null;
          if (example === void 0) {
            example = null;
            type = schema.type;
            if (Array.isArray(type) && schema.type.length > 0) {
              type = schema.type[0];
            }
            if (!type) {
              type = (0, _infer.inferType)(schema);
            }
            let sampler = _openapiSampler._samplers[type];
            if (sampler) {
              example = sampler(schema, options, spec, context);
            }
          }
          (0, _utils.popSchemaStack)(seenSchemasStack, context);
          return {
            value: example,
            readOnly: schema.readOnly,
            writeOnly: schema.writeOnly,
            type
          };
          function traverseOneOrAnyOf(schema2, selectedSubSchema) {
            const inferred = tryInferExample(schema2);
            if (inferred !== void 0) {
              return inferred;
            }
            const localExample = traverse(_objectSpread(_objectSpread({}, schema2), {}, {
              oneOf: void 0,
              anyOf: void 0
            }), options, spec, context);
            const subSchemaExample = traverse(selectedSubSchema, options, spec, context);
            if (typeof localExample.value === "object" && typeof subSchemaExample.value === "object") {
              const mergedExample = (0, _utils.mergeDeep)(localExample.value, subSchemaExample.value);
              return _objectSpread(_objectSpread({}, subSchemaExample), {}, {
                value: mergedExample
              });
            }
            return subSchemaExample;
          }
        }
      }, { "./allOf": 4, "./infer": 5, "./openapi-sampler": 6, "./utils": 15, "json-pointer": 3 }], 15: [function(require2, module3, exports3) {
        "use strict";
        Object.defineProperty(exports3, "__esModule", {
          value: true
        });
        exports3.SKIP_SYMBOL = void 0;
        exports3.applyXMLAttributes = applyXMLAttributes;
        exports3.ensureMinLength = ensureMinLength;
        exports3.getResultForCircular = getResultForCircular;
        exports3.getXMLAttributes = getXMLAttributes;
        exports3.mergeDeep = mergeDeep;
        exports3.popSchemaStack = popSchemaStack;
        exports3.toRFCDateTime = toRFCDateTime;
        exports3.uuid = uuid;
        function ownKeys(e, r) {
          var t = Object.keys(e);
          if (Object.getOwnPropertySymbols) {
            var o = Object.getOwnPropertySymbols(e);
            r && (o = o.filter(function(r2) {
              return Object.getOwnPropertyDescriptor(e, r2).enumerable;
            })), t.push.apply(t, o);
          }
          return t;
        }
        function _objectSpread(e) {
          for (var r = 1; r < arguments.length; r++) {
            var t = null != arguments[r] ? arguments[r] : {};
            r % 2 ? ownKeys(Object(t), true).forEach(function(r2) {
              _defineProperty(e, r2, t[r2]);
            }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(e, Object.getOwnPropertyDescriptors(t)) : ownKeys(Object(t)).forEach(function(r2) {
              Object.defineProperty(e, r2, Object.getOwnPropertyDescriptor(t, r2));
            });
          }
          return e;
        }
        function _defineProperty(e, r, t) {
          return (r = _toPropertyKey(r)) in e ? Object.defineProperty(e, r, { value: t, enumerable: true, configurable: true, writable: true }) : e[r] = t, e;
        }
        function _toPropertyKey(t) {
          var i = _toPrimitive(t, "string");
          return "symbol" == typeof i ? i : i + "";
        }
        function _toPrimitive(t, r) {
          if ("object" != typeof t || !t) return t;
          var e = t[Symbol.toPrimitive];
          if (void 0 !== e) {
            var i = e.call(t, r || "default");
            if ("object" != typeof i) return i;
            throw new TypeError("@@toPrimitive must return a primitive value.");
          }
          return ("string" === r ? String : Number)(t);
        }
        const SKIP_SYMBOL = exports3.SKIP_SYMBOL = /* @__PURE__ */ Symbol("skip");
        function pad(number) {
          if (number < 10) {
            return "0" + number;
          }
          return number;
        }
        function toRFCDateTime(date, omitTime, omitDate, milliseconds) {
          var res = omitDate ? "" : date.getUTCFullYear() + "-" + pad(date.getUTCMonth() + 1) + "-" + pad(date.getUTCDate());
          if (!omitTime) {
            res += "T" + pad(date.getUTCHours()) + ":" + pad(date.getUTCMinutes()) + ":" + pad(date.getUTCSeconds()) + (milliseconds ? "." + (date.getUTCMilliseconds() / 1e3).toFixed(3).slice(2, 5) : "") + "Z";
          }
          return res;
        }
        ;
        function ensureMinLength(sample, min) {
          if (min > sample.length) {
            return sample.repeat(Math.trunc(min / sample.length) + 1).substring(0, min);
          }
          return sample;
        }
        function mergeDeep() {
          const isObject = (obj) => obj && typeof obj === "object";
          for (var _len = arguments.length, objects = new Array(_len), _key = 0; _key < _len; _key++) {
            objects[_key] = arguments[_key];
          }
          return objects.reduce((prev, obj) => {
            Object.keys(obj || {}).forEach((key) => {
              const pVal = prev[key];
              const oVal = obj[key];
              if (isObject(pVal) && isObject(oVal)) {
                prev[key] = mergeDeep(pVal, oVal);
              } else {
                prev[key] = oVal;
              }
            });
            return prev;
          }, Array.isArray(objects[objects.length - 1]) ? [] : {});
        }
        function uuid(str) {
          var hash = hashCode(str);
          var random = jsf32(hash, hash, hash, hash);
          var uuid2 = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, (c) => {
            var r = random() * 16 % 16 | 0;
            return (c == "x" ? r : r & 3 | 8).toString(16);
          });
          return uuid2;
        }
        function getResultForCircular(type) {
          return {
            value: type === "object" ? {} : type === "array" ? [] : void 0
          };
        }
        function popSchemaStack(seenSchemasStack, context) {
          if (context) seenSchemasStack.pop();
        }
        function getXMLAttributes(schema) {
          var _schema$xml, _schema$xml2, _schema$xml3, _schema$xml$attribute, _schema$xml4, _schema$xml$wrapped, _schema$xml5;
          return {
            name: (schema === null || schema === void 0 || (_schema$xml = schema.xml) === null || _schema$xml === void 0 ? void 0 : _schema$xml.name) || "",
            prefix: (schema === null || schema === void 0 || (_schema$xml2 = schema.xml) === null || _schema$xml2 === void 0 ? void 0 : _schema$xml2.prefix) || "",
            namespace: (schema === null || schema === void 0 || (_schema$xml3 = schema.xml) === null || _schema$xml3 === void 0 ? void 0 : _schema$xml3.namespace) || null,
            attribute: (_schema$xml$attribute = schema === null || schema === void 0 || (_schema$xml4 = schema.xml) === null || _schema$xml4 === void 0 ? void 0 : _schema$xml4.attribute) !== null && _schema$xml$attribute !== void 0 ? _schema$xml$attribute : false,
            wrapped: (_schema$xml$wrapped = schema === null || schema === void 0 || (_schema$xml5 = schema.xml) === null || _schema$xml5 === void 0 ? void 0 : _schema$xml5.wrapped) !== null && _schema$xml$wrapped !== void 0 ? _schema$xml$wrapped : false
          };
        }
        function resolveNodeType(schema) {
          const xml = schema === null || schema === void 0 ? void 0 : schema.xml;
          if (xml !== null && xml !== void 0 && xml.nodeType) {
            return xml.nodeType;
          }
          if ((xml === null || xml === void 0 ? void 0 : xml.attribute) === true) {
            return "attribute";
          }
          if ((xml === null || xml === void 0 ? void 0 : xml.wrapped) === true && (schema === null || schema === void 0 ? void 0 : schema.type) === "array") {
            return "element";
          }
          if (schema !== null && schema !== void 0 && schema.$ref || schema !== null && schema !== void 0 && schema.$dynamicRef || (schema === null || schema === void 0 ? void 0 : schema.type) === "array") {
            return "none";
          }
          if (schema !== null && schema !== void 0 && schema.oneOf || schema !== null && schema !== void 0 && schema.anyOf || schema !== null && schema !== void 0 && schema.allOf) {
            return "none";
          }
          return "element";
        }
        function applyXMLAttributes(result) {
          let schema = arguments.length > 1 && arguments[1] !== void 0 ? arguments[1] : {};
          let context = arguments.length > 2 && arguments[2] !== void 0 ? arguments[2] : {};
          const {
            value: oldValue
          } = result;
          const {
            propertyName: oldPropertyName
          } = context;
          const {
            name,
            prefix,
            namespace
          } = getXMLAttributes(schema);
          const effectiveNodeType = resolveNodeType(schema);
          let propertyName = name || oldPropertyName ? "".concat(prefix ? prefix + ":" : "").concat(name || oldPropertyName) : null;
          let value = typeof oldValue === "object" ? Array.isArray(oldValue) ? [...oldValue] : _objectSpread({}, oldValue) : oldValue;
          switch (effectiveNodeType) {
            case "attribute":
              if (propertyName) {
                propertyName = "$".concat(propertyName);
              }
              break;
            case "text":
              propertyName = "#text";
              break;
            case "cdata":
              propertyName = "#cdata";
              break;
            case "none":
              if (schema.type === "array") {
                propertyName = null;
                if (schema.example !== void 0) {
                  var _schema$items;
                  propertyName = ((_schema$items = schema.items) === null || _schema$items === void 0 || (_schema$items = _schema$items.xml) === null || _schema$items === void 0 ? void 0 : _schema$items.name) || propertyName;
                }
              } else {
                propertyName = null;
              }
              break;
            default:
              if (schema.type === "array") {
                if (Array.isArray(value)) {
                  value = {
                    [propertyName]: [...value]
                  };
                }
              }
              break;
          }
          if (namespace && effectiveNodeType !== "text" && effectiveNodeType !== "cdata" && effectiveNodeType !== "none") {
            if (typeof value === "object") {
              value["$xmlns".concat(prefix ? ":" + prefix : "")] = namespace;
            } else {
              value = {
                ["$xmlns".concat(prefix ? ":" + prefix : "")]: namespace,
                ["#text"]: value
              };
            }
          }
          return {
            propertyName,
            value
          };
        }
        function hashCode(str) {
          var hash = 0;
          if (str.length == 0) return hash;
          for (var i = 0; i < str.length; i++) {
            var char = str.charCodeAt(i);
            hash = (hash << 5) - hash + char;
            hash = hash & hash;
          }
          return hash;
        }
        function jsf32(a, b, c, d) {
          return function() {
            a |= 0;
            b |= 0;
            c |= 0;
            d |= 0;
            var t = a - (b << 27 | b >>> 5) | 0;
            a = b ^ (c << 17 | c >>> 15);
            b = c + d | 0;
            c = d + t | 0;
            d = a + t | 0;
            return (d >>> 0) / 4294967296;
          };
        }
      }, {}] }, {}, [6])(6);
    });
  }
});

export {
  require_json_pointer,
  require_openapi_sampler
};
