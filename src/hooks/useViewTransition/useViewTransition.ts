import {useEffect, useState} from "react"
import styles from "./styles.module.scss"


type Props = {
    currentView: string

    viewList: [{
        id: string
        element: HTMLElement
    }]

    duration?: number // TODO
    transition?: "slideX" | "slideY" | "rotate-left" | "rotate-right" | "flipX" | "flipY" // TODO
    timingFunction?: string // TODO

}


export default function useViewTransition({currentView, viewList}: Props) {
    const [firstRender, setFirstRender] = useState(true)
    const [lastView, setLastView] = useState<string | null>(null)
    useEffect(() => {
        if (firstRender) {
            setFirstRender(false)
            setLastView(currentView)
            return
        }
        viewList.forEach(({id, element}) => {
            if (id === currentView) {
                show(element)
                setLastView(currentView)
            } else {
                hide(element)
            }
        })
    }, [currentView])

    function hide(element: HTMLElement) {
        element.classList.remove(styles.active)
        element.classList.add(styles.hidden)
        //TODO: transition
    }

    function show(element: HTMLElement) {
        element.classList.remove(styles.hidden)
        element.classList.add(styles.active)
        //TODO: transition
    }

}