"use strict";
const SERVER_ID = 0x01;
const APP_ID = 0x02;
const SESSION = 0x03;
const CRYPTO = 0x04;
const DEVICE_ID = 0x05;
const DEVICE_NM = 0x08;
const PROXY_ID = 0x06;
const LISTEN = 0x07;
const CLIENT_ID = 0x09;
const EGG_TYPE_PING = 0x00;
const EGG_TYPE_PANG = 0xFF;
const EGG_TYPE_SESSION = 0xFE;
const EGG_TYPE_CLOSE = 0xFD;
const EGG_TYPE_OAUTH = 0xFC;
const EGG_TYPE_REGISTER = 0xFB;
const EGG_TYPE_ROGER = 0xFA;
const EGG_TYPE_ERROR = 0xF9;
const EGG_TYPE_STRING = 0xEF;
const EGG_TYPE_EVENT_NAME = 0xEF;
const EGG_TYPE_INT32 = 0xEE;
const EGG_TYPE_INT64 = 0xED;
const EGG_TYPE_UINT32 = 0xEC;
const EGG_TYPE_UINT64 = 0xEB;
const EGG_TYPE_JSON = 0xEA;
const EGG_TYPE_EVENT_ARGUMENTS = 0xEA;
const EGG_TYPE_BIN = 0xE9;
const EGG_TYPE_INT = 0xE8;
const EGG_TYPE_UINT = 0xE7;
const EGG_TYPE_COMMAND = 0xD7;
const EGG_TYPE_RESPONSE = 0xD6;
// 打印服务
const EGG_TYPE_PRINTER = 0x11;
const EGG_TYPE_ORDER = 0x12;
const EGG_TYPE_TODO = 0x13
const EGG_TYPE_MESSAGE = 0x14

const EGG_TYPE_SHAKE = 0x0F;
const EGG_TYPE_HEADERS = 0x0E;
const EGG_TYPE_EVENT_MAP = 0x0D;
const EGG_TYPE_EVENT_LIST = 0x0C;
const EGG_TYPE_EVENT_BACK_MAP = 0x0B;
const EGG_TYPE_CRYPTO = 0x01;
const EGG_TYPE_VALUES = 0x02;
const EGG_HEADER_LENGTH = 9;
const EGG_PROTOCOL = "TLAV";
class Egg {
    constructor(type, ai, value) {
        this.Type = type;
        this.Length = value.length;
        this.Ai = ai;
        this.Value = value;
        this.pb = "TLAV";
    }
    Encode() {
        return new Uint8Array(this.Bytes());
    }
    Bytes() {
        const buffer = [];
        buffer.push(...StringToBytes(this.pb));
        buffer.push(...uint8ToBytes(this.Type));
        buffer.push(...uint24ToBytes(this.Length));
        buffer.push(...uint8ToBytes(this.Type));
        buffer.push(...this.Value);
        return buffer;
    }
    Size() {
        return this.Length + EGG_HEADER_LENGTH;
    }
    MustValues() {
        const vs = [];
        let buf = [...this.Value];
        do {
            const b = readEgg(new Uint8Array(buf));
            if (b == null) {
                break;
            }
            vs.push(b.Value);
            buf = buf.slice(b.Size());
        } while (buf.length > 0);
        return vs;
    }
    MustEventMap() {
        let name = null;
        let rst = {};
        if (this.Type == EGG_TYPE_EVENT_MAP || this.Type == EGG_TYPE_EVENT_BACK_MAP) {
            let buf = this.Value.slice(0);
            const n = readEggWithType(buf, EGG_TYPE_EVENT_NAME);
            if (n != null) {
                buf = buf.slice(n.Size());
                const arg = readEggWithType(buf, EGG_TYPE_EVENT_ARGUMENTS);
                if (arg != null) {
                    name = BytestoString(n.Value);
                    const evts = arg.MustValues();
                    rst = deMapArgments(evts);
                }
            }
        }
        return [name, rst];
    }
}
function deMapArgments(params) {
    const rst = {};
    if (params.length > 1) {
        const name = BytestoString(params[0]);
        const keys = name.split(",");
        for (let i = 0; i < keys.length; i = i + 1) {
            const k = keys[i];
            rst[k] = params[i + 1];
        }
    }
    return rst;
}
function readEggValue(buffer) {
    if (buffer.length < EGG_HEADER_LENGTH) {
        return null;
    }
    const pb = buffer.slice(0, 4);
    const plus = pb.length;
    const type = buffer[0 + plus];
    const lenBuf = buffer.slice(1 + plus, 4 + plus);
    const bobyLength = bytesToUint24(lenBuf);
    const eggLength = bobyLength + EGG_HEADER_LENGTH;
    if (buffer.length < eggLength) {
        return null;
    }
    return buffer.slice(EGG_HEADER_LENGTH, eggLength + EGG_HEADER_LENGTH);
}
;
function readAEggData(buffer) {
    if (buffer.length < EGG_HEADER_LENGTH) {
        return null;
    }
    const pb = buffer.slice(0, 4);
    const plus = pb.length;
    const type = buffer[0 + plus];
    const lenBuf = buffer.slice(1 + plus, 4 + plus);
    const bobyLength = bytesToUint24(lenBuf);
    const eggLength = bobyLength + EGG_HEADER_LENGTH;
    if (buffer.length < eggLength) {
        return null;
    }
    return buffer.slice(0, eggLength);
}
;
function readEgg(buffer) {
    let receive = [...buffer];
    if (receive.length < EGG_HEADER_LENGTH) {
        return null;
    }
    const pb = receive.slice(0, 4);
    if (BytestoString(pb) !== EGG_PROTOCOL) {
        return null;
    }
    const plus = pb.length;
    const type = receive[0 + plus];
    const lenBuf = receive.slice(1 + plus, 4 + plus);
    const bobyLength = bytesToUint24(lenBuf);
    if (receive.length < bobyLength + EGG_HEADER_LENGTH) {
        return null;
    }
    const ai = receive[4 + plus];
    var bytes = receive.slice(EGG_HEADER_LENGTH, bobyLength + EGG_HEADER_LENGTH);
    return new Egg(type, ai, bytes);
}
;
function readEggWithType(buffer, t) {
    const buf = buffer.slice(0);
    if (buf.length < EGG_HEADER_LENGTH) {
        return null;
    }
    const pb = buf.slice(0, 4);
    if (BytestoString(pb) !== EGG_PROTOCOL) {
        return null;
    }
    const plus = pb.length;
    const type = buf[0 + plus];
    if (type != t) {
        return null;
    }
    const lenBuf = buf.slice(1 + plus, 4 + plus);
    const bobyLength = bytesToUint24(lenBuf);
    if (buf.length < bobyLength + EGG_HEADER_LENGTH) {
        return null;
    }
    const ai = buf[4 + plus];
    const value = buf.slice(EGG_HEADER_LENGTH, bobyLength + EGG_HEADER_LENGTH);
    return new Egg(type, ai, value);
}
;
function FromString(str, t = EGG_TYPE_STRING) {
    const buffer = StringToBytes(str);
    return new Egg(t, 0, buffer);
}
function FromPong() {
    return new Egg(EGG_TYPE_PANG, 0, uint8ToBytes(EGG_TYPE_PANG));
}
function FromPing() {
    return new Egg(EGG_TYPE_PING, 0, uint8ToBytes(EGG_TYPE_PING));
}
// Type and Values
function FromTValues(obj) {
    const t = EGG_TYPE_VALUES;
    const buffer = [];
    const ai = 0;
    for (const key in obj) {
        if (Object.hasOwnProperty.call(obj, key)) {
            const element = obj[key];
            const data = StringToBytes(element);
            const item = [];
            item.push(...StringToBytes('TLAV'));
            item.push(...uint8ToBytes(Number(key)));
            item.push(...uint24ToBytes(data.length));
            item.push(...uint8ToBytes(ai));
            item.push(...data);
            buffer.push(...item);
        }
    }
    return new Egg(t, ai, buffer);
}
function FromMap(args) {
    const t = EGG_TYPE_EVENT_ARGUMENTS;
    const buffer = [];
    const kvs = enMapArgments(args);
    for (let i = 0; i < kvs.length; i = i + 1) {
        const data = kvs[i];
        const item = [];
        item.push(...StringToBytes('TLAV'));
        item.push(...uint8ToBytes(EGG_TYPE_BIN));
        item.push(...uint24ToBytes(data.length));
        item.push(...uint8ToBytes(i));
        item.push(...data);
        buffer.push(...item);
    }
    return new Egg(t, 0, buffer);
}
function enMapArgments(args) {
    const rst = {};
    const values = [];
    const names = [];
    values.push([]);
    for (const key in args) {
        if (Object.hasOwnProperty.call(args, key)) {
            const value = args[key];
            names.push(key);
            values.push(value);
        }
    }
    if (names.length + 1 == values.length) {
        values[0] = StringToBytes(names.join(','));
    }
    return values;
}
function FromEventMap(event, data) {
    const buffer = [];
    const name = FromString(event);
    name.Ai = 0;
    buffer.push(...name.Bytes());
    const params = FromMap(data);
    params.Ai = 1;
    buffer.push(...params.Bytes());
    let tp = EGG_TYPE_EVENT_MAP;
    if (event[0] === ":") {
        tp = EGG_TYPE_EVENT_BACK_MAP;
    }
    return new Egg(tp, 0, buffer);
}
function StringToBytes(str) {
    const buffer = [];
    for (let i = 0; i < str.length; i++) {
        const c = str.charCodeAt(i);
        buffer.push(...GetCodeBytes(c));
    }
    return buffer;
}
function BytestoString(buffer) {
    let utf8decoder = new TextDecoder();
    return utf8decoder.decode(new Uint8Array(buffer));
}
function uint24ToBytes(v) {
    v = v & 0x00ffffff;
    const b = [];
    b[0] = (v & 0xFF);
    b[1] = ((v >> 8) & 0xFF);
    b[2] = ((v >> 16) & 0xFF);
    return b;
}
// big
function bytesToUint24(v) {
    let uint1 = v[0];
    let uint2 = v[1];
    let uint3 = v[2];
    return (uint3 << 16) | (uint2 << 8) | uint1;
}
function uint8ToBytes(v) {
    const b = [];
    b[0] = (v & 0xFF);
    return b;
}
function GetCodeBytes(c) {
    const buffer = [];
    if (c <= 0x7F) {
        buffer.push(c);
    }
    else if (c <= 0xFF) {
        buffer.push((c >> 6) | 0xC0);
        buffer.push((c & 0x3F) | 0x80);
    }
    else if (c <= 0xFFFF) {
        buffer.push((c >> 12) | 0xE0);
        buffer.push(((c >> 6) & 0x3F) | 0x80);
        buffer.push((c & 0x3F) | 0x80);
    }
    else if (c <= 0x1FFFFF) {
        buffer.push((c >> 18) | 0xF0);
        buffer.push(((c >> 12) & 0x3F) | 0x80);
        buffer.push(((c >> 6) & 0x3F) | 0x80);
        buffer.push((c & 0x3F) | 0x80);
    }
    else if (c <= 0x3FFFFFF) { //后面两种情况一般不大接触到，看了下protobuf.js中的utf8，他没去实现
        buffer.push((c >> 24) | 0xF8);
        buffer.push(((c >> 18) & 0x3F) | 0x80);
        buffer.push(((c >> 12) & 0x3F) | 0x80);
        buffer.push(((c >> 6) & 0x3F) | 0x80);
        buffer.push((c & 0x3F) | 0x80);
    }
    else { //Math.pow(2, 32) - 1
        buffer.push((c >> 30) & 0x1 | 0xFC);
        buffer.push(((c >> 24) & 0x3F) | 0x80);
        buffer.push(((c >> 18) & 0x3F) | 0x80);
        buffer.push(((c >> 12) & 0x3F) | 0x80);
        buffer.push(((c >> 6) & 0x3F) | 0x80);
        buffer.push((c & 0x3F) | 0x80);
    }
    return buffer;
}
function getValue(k) {
    switch (k) {
        case SERVER_ID:
            return "0";
        case CLIENT_ID:
            return "123";
        case CRYPTO:
            return "11";
        case SESSION:
            return "1245";
        case APP_ID:
            return "a1254";
        case DEVICE_ID:
            return "abdfdef";
        case DEVICE_NM:
            return "dv_name";
        case PROXY_ID:
            return "47";
        case LISTEN:
            return "echo";
        default:
            return "";
    }
}
class Listeners {
    constructor() {
        this.items = {};
    }
    has(key) {
        return this.items.hasOwnProperty(key);
    }
    set(key, val) {
        const rst = this.get(key);
        if (rst === null) {
            this.items[key] = [val];
        }
        else if (Array.isArray(rst)) {
            this.items[key] = rst.concat(val);
        }
    }
    delete(key) {
        if (this.has(key)) {
            delete this.items[key];
        }
        return false;
    }
    get(key) {
        return this.has(key) ? this.items[key] : null;
    }
}
class SSTWSock {
    constructor(uri) {
        this.uri = uri;
        this.binaryType = "arraybuffer";
        this.linked = false;
        this.authed = false;
        this.keepLive = true;
        // 这个只能声明一次
        this.listener = new Listeners();
        this.buffer = [];
        this.websock = null;
        this.waitTime = 0;
        this.keepTimer = undefined;
        this.status = 0;
        this.stoped = false;
    }
    init() {
        if (this.keepTimer != undefined) {
            clearInterval(this.keepTimer);
        }
        this.linked = false;
        this.authed = false;
        this.keepLive = true;
        this.buffer = [];
        this.websock = null;
        this.waitTime = 0;
        this.stoped = false;
        this.keepTimer = setInterval(() => {
            this.waitTime += 2500;
            if (this.waitTime > 60000) {
                this.waitTime = 0;
                this.ping();
            }
        }, 2500);
    }
    OnHeader(f) {
        this.onHeaderFunc = f;
    }
    OnData(f) {
        this.onDataFunc = f;
    }
    OnClose(f) {
        this.onCloseFunc = f;
    }
    OnOpen(f) {
        this.onOpenFunc = f;
    }
    OnReady(f) {
        this.onReadyFunc = f;
    }
    OnError(f) {
        this.onErrorFunc = f;
    }
    ping() {
        console.log("ping");
        const eg = FromPing();
        this.Send(eg.Encode());
    }
    Connect() {
        this.init();
        const ws = new WebSocket(this.uri);
        ws.binaryType = this.binaryType;
        ws.onmessage = (evt) => {
            if (evt.data instanceof ArrayBuffer) {
                let data = new Uint8Array(evt.data);
                if (!this.authed && data.length == 2) {
                    if (data[0] == 0xFF && data[1] == 0x00) {
                        this.authed = true;
                        this.status = 1;
                        this.buffer = [];
                        if (this.onReadyFunc)
                            this.onReadyFunc(true);
                        return;
                    }
                }
                this.buffer = this.buffer.concat(...data);
                const eg = readEgg(data);
                if (eg != null) {
                    this.buffer = this.buffer.slice(eg.Size());
                    if (eg.Type == EGG_TYPE_HEADERS) {
                        const header = {};
                        const v = [...eg.Value];
                        if (this.onHeaderFunc != undefined) {
                            for (let i = 0; i < v.length; i++) {
                                const p = v[i];
                                header[p] = this.onHeaderFunc(p);
                            }
                        }
                        const h = FromTValues(header);
                        ws.send(h.Encode());
                    }
                    else if (eg.Type === EGG_TYPE_EVENT_BACK_MAP || eg.Type === EGG_TYPE_EVENT_MAP) {
                        const [evt, args] = eg.MustEventMap();
                        if (evt != null) {
                            const cbs = this.listener.get(evt);
                            if (cbs != null) {
                                cbs.map((res) => {
                                    res.call(null, args, eg);
                                });
                                // 临时监听才需要删
                                if (evt[0] === ":")
                                    this.listener.delete(evt);
                            }
                        }
                    }
                    else if (eg.Type === EGG_TYPE_PANG) {
                    }
                    else if (eg.Type === EGG_TYPE_PING) {
                        const eg = FromPong();
                        this.Send(eg.Encode());
                    }
                    else {
                        if (this.onDataFunc != undefined)
                            this.onDataFunc(eg);
                    }
                }
            }
        };
        ws.onopen = (evt) => {
            if (this.onOpenFunc != undefined)
                this.onOpenFunc(evt);
            if (ws.readyState === WebSocket.OPEN) {
                ws.send(new Uint8Array([0x00, 0xFF]));
            }
        };
        ws.onclose = evt => {
            if (this.onCloseFunc != undefined) {
                this.onCloseFunc(evt);
            }
            if (this.keepTimer != undefined) {
                clearInterval(this.keepTimer);
            }
            this.linked = false;
            this.authed = false;
            if (this.status < 0)
                return;
            this.status = -1;
            if (this.keepLive == true && this.stoped == false) {
                const r = (Math.round(Math.random() * 100) + 10) * 1000;
                setTimeout(() => {
                    this.status = 0;
                    this.Connect();
                }, r);
                console.log(r / 1000, "后重新链接");
            }
        };
        ws.onerror = (evt) => {
            if (this.onErrorFunc != undefined)
                this.onErrorFunc(evt);
        };
        this.websock = ws;
        this.linked = true;
    }
    Emit(event, kv, callback) {
        if (!this.authed) {
            console.log("没有");
            return;
        }
        if (callback != undefined) {
            const r = this.randStr();
            const ck = `:${event}_${r}`;
            this.listener.set(ck, callback);
            kv['__cbk'] = StringToBytes(ck);
        }
        const eg = FromEventMap(event, kv);
        this.Send(eg.Encode());
    }
    EmitCallBack(eg, kv) {
        const [evt, args] = eg.MustEventMap();
        if (evt != null) {
            const k = "__cbk";
            if (args.hasOwnProperty(k)) {
                const cb = BytestoString(args["__cbk"]);
                const e = FromEventMap(cb, kv);
                this.Send(e.Encode());
            }
        }
    }
    AddEvent(event, callback) {
        this.listener.set(event, callback);
    }
    RemoveEvent(event) {
        this.listener.delete(event);
    }
    Send(data) {
        if (this.linked === false) {
            this.Connect();
        }
        const ws = this.websock;
        if (ws != null) {
            if (ws.readyState === WebSocket.OPEN) {
                ws.send(data);
                return null;
            }
        }
        return "已断开链接";
    }
    Close(code, reason) {
        if (this.keepTimer != undefined) {
            clearInterval(this.keepTimer);
        }
        const ws = this.websock;
        if (ws != null) {
            if (ws.readyState !== WebSocket.CLOSED && ws.readyState !== WebSocket.CLOSING) {
                ws.close(code, reason);
            }
        }
    }
    Stop() {
        this.stoped = true;
        this.Close();
    }
    randStr() {
        const n = 8;
        let result = '';
        const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
        const charactersLength = characters.length;
        for (let i = 0; i < n; i++) {
            result += characters.charAt(Math.floor(Math.random() * charactersLength));
        }
        return result;
    }
}
// const msg: Value = [239, 11, 0, 0, 0, 229, 136, 152, 88, 228, 184, 141, 231, 137, 155, 66]
// const egg = readEgg(msg);
// console.log(egg?.Value)
// if (egg != null) {
//     console.log(BytestoString(egg?.Value))
// }
// tsc .\egg.ts
