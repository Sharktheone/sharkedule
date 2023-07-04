import {useEffect, useState} from "react"
import styles from "./styles.module.scss"


export type viewRef = {
    id: string
    element: HTMLElement
}


export default function useViewTransition(currentView: string, viewList: viewRef[], duration?: number, transition?: string, timingFunction?: string) {
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
        element.classList.add(styles.hidden, styles[transition ?? ""])
    }

    function show(element: HTMLElement) {
        element.classList.remove(styles.hidden)
        element.classList.add(styles.active, styles[transition ?? ""])
        //TODO: transition
    }

}