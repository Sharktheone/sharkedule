import {useEffect} from "react"
import styles from "./styles.module.scss"


export type viewRef = {
    id: string
    element: HTMLElement
}

//TODO
// initial state: show currentView / default
// onChange show newCurrentView -> transition to new -> hide old
//
//
/// # Change
/// - 1. showNew
///     - 1.1. add a show class, but  it needs to be hidden in overflow => translate to right or left (depending on the index)
///     - 1.2. add a transition class
///         - 1.2.1. translate old and new to left or right (depending on the index)
///         - 1.2.2. remove transition class
/// - 2. hideOld



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