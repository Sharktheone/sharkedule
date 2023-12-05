import {MutableRefObject, useRef} from "react"


export function useClickOutside<T extends HTMLElement = any>(handler: () => void, events: string[] = ["mousedown", "touchstart"], nodes?: (HTMLElement | null)[]) {
    //TODO


    return useRef<T>()
}