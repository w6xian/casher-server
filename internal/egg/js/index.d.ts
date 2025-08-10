interface Egg {
    pb: string
    Type: number;
    Length: number;
    Ai: number;
    Value: number[];
}
interface Listeners {
    items: { [key: string]: Function[] };
}

interface Sock {
    websock: WebSocket
    uri: string
    binaryType: BinaryType
    linked: boolean
    buffer: number[]
    wLock: boolean
    nIntervId: number | undefined
    onHeaderFunc: FunctionUndefined
    onDataFunc: FunctionUndefined
    onOpenFunc: FunctionUndefined
    onCloseFunc: FunctionUndefined
    onErrorFunc: FunctionUndefined
    keepLive: boolean
    authed: boolean
    listener: Listeners
}

type Dictionary = { [key: string]: Function[] }
type ListenFunc = Function[] | null
type EggNull = Egg | null
type Value = number[]
type ValueNull = number[] | null
type StringNull = string | null
type FunctionUndefined = Function | undefined
type NumberUndefined = number | undefined
type Arguments = { [key: string]: Value }