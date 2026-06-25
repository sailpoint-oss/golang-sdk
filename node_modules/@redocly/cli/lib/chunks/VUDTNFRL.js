import { createRequire as __createRequire } from 'node:module';
const require = __createRequire(import.meta.url);
import {
  require_execAsync
} from "./AIT3U3JT.js";
import {
  esm_exports,
  init_esm
} from "./XHPY6UJF.js";
import {
  __commonJS,
  __require,
  __toCommonJS
} from "./5ILQMFXK.js";

// ../../node_modules/@opentelemetry/resources/build/src/detectors/platform/node/machine-id/getMachineId-bsd.js
var require_getMachineId_bsd = __commonJS({
  "../../node_modules/@opentelemetry/resources/build/src/detectors/platform/node/machine-id/getMachineId-bsd.js"(exports) {
    Object.defineProperty(exports, "__esModule", { value: true });
    exports.getMachineId = void 0;
    var fs_1 = __require("fs");
    var execAsync_1 = require_execAsync();
    var api_1 = (init_esm(), __toCommonJS(esm_exports));
    async function getMachineId() {
      try {
        const result = await fs_1.promises.readFile("/etc/hostid", { encoding: "utf8" });
        return result.trim();
      } catch (e) {
        api_1.diag.debug(`error reading machine id: ${e}`);
      }
      try {
        const result = await (0, execAsync_1.execAsync)("kenv -q smbios.system.uuid");
        return result.stdout.trim();
      } catch (e) {
        api_1.diag.debug(`error reading machine id: ${e}`);
      }
      return void 0;
    }
    exports.getMachineId = getMachineId;
  }
});
export default require_getMachineId_bsd();
