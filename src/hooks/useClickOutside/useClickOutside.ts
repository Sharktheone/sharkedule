import {useEffect, useRef} from "react"


export function useClickOutside<T extends HTMLElement = any>(handler: () => void, nodes?: (HTMLElement | null)[], events: string[] = ["mousedown", "touchstart"]) {

    const ref = useRef<T>()

    useEffect(() => {
        const listener = (event: Event) => {
            const target = event.target as HTMLElement

            if (ref.current && !ref.current.contains(target)) {
                handler()
            }
        }

        events.forEach(event => {
            document.addEventListener(event, listener)
        })

        return () => {
            events.forEach(event => {
                document.removeEventListener(event, listener)
            })
        }
    }, [handler, events, ref, nodes])


    return ref
}