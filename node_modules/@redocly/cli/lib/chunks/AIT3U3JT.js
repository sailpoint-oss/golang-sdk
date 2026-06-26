import { createRequire as __createRequire } from 'node:module';
const require = __createRequire(import.meta.url);
import {
  __commonJS,
  __require
} from "./5ILQMFXK.js";

// ../../node_modules/@opentelemetry/resources/build/src/detectors/platform/node/machine-id/execAsync.js
var require_execAsync = __commonJS({
  "../../node_modules/@opentelemetry/resources/build/src/detectors/platform/node/machine-id/execAsync.js"(exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.execAsync = void 0;
    var child_process = __require("child_process");
    var util = __require("util");
    exports.execAsync = util.promisify(child_process.exec);
  }
});

export {
  require_execAsync
};
