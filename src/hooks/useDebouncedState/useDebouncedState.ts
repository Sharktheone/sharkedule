import {useRef, useState} from "react"


export function useDebouncedState<S = undefined>(initial: S, delay: 1000) {

    const [state, setState] = useState<S>(initial)
    const timeout = useRef<number | null>(null)

    function debounce(val: S) {

        if (timeout.current) clearTimeout(timeout.current)

        //@ts-ignore
        timeout.current = setTimeout(() => {
            setState(val)
            timeout.current = null
        }, delay)
    }

    return [state, debounce] as const
}