import {useEffect, useState} from "react"
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



export default function useViewTransition(currentView: string, viewList: viewRef[], duration: number = 500, transition?: string, timingFunction?: string) {
    const [lastView, setLastView] = useState<string>(currentView)


    useEffect(() => {
        viewList.forEach(({id, element}) => {
            if (id === currentView) {
                if (!element) return
                show(element)
                // setLastView(currentView)
            } else if (id !== lastView) {
                if (!element) return
                hide(element)
            }
        })
        setLastView(currentView)
    }, [currentView, viewList])


    function direction(currentView: string, newView: string) {
        const currentIndex = viewList.findIndex(({id}) => id === currentView)
        const newIndex = viewList.findIndex(({id}) => id === newView)

        if (currentIndex < newIndex) {
            return styles.hiddenRight
        } else {
            return styles.hiddenLeft
        }
    }

    function getOldElement() {
        return viewList.find(({id}) => id === lastView)?.element
    }

    function hide(element: HTMLElement) {
        element.classList.remove(styles.active)
        element.classList.add(styles.hidden, styles[transition ?? ""])
    }

    function show(element: HTMLElement) {
        const old = getOldElement()
        element.classList.remove(styles.hidden)
        if (!old) {
            element.classList.add(styles.active)
            return
        }
        element.classList.add(direction(currentView, lastView), styles[transition ?? ""])
        old.classList.add(direction(lastView, currentView), styles[transition ?? ""])
        setTimeout(() => {
            element.classList.remove(direction(currentView, lastView), styles[transition ?? ""])
            old.classList.remove(direction(lastView, currentView), styles[transition ?? ""])
            element.classList.add(styles.active)
        }, duration)
        //TODO: transition
    }
}