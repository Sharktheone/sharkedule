import {useEffect} from "react"
import styles from "./styles.module.scss"


export type viewRef = {
    id: string
    element: HTMLElement
}

export default function useViewTransition(currentView: string, viewList: viewRef[], duration?: number, transition?: string, timingFunction?: string) {
    useEffect(() => {
        viewList.forEach(({id, element}) => {
            console.log(element, id)
            if (id === currentView) {
                if (!element) return
                show(element)
                // setLastView(currentView)
            } else {
                if (!element) return
                hide(element)
            }
        })
    }, [currentView, viewList])

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