import {useEffect} from "react"
import styles from "./styles.module.scss"


export type viewRef = {
    id: string
    element: HTMLElement
}


export default function useViewTransition(currentView: string, lastView: string, viewList: viewRef[], duration: number = 10000, transition?: string, timingFunction?: string) {
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
        element.classList.add(styles.hidden)
    }

    function show(element: HTMLElement) {
        const old = getOldElement()
        element.classList.remove(styles.hidden)
        if (!old) {
            element.classList.add(styles.active)
            return
        }

        console.log(direction(currentView, lastView))
        element.classList.add(direction(lastView, currentView), styles.slideX)
        old.classList.add(direction(currentView, lastView), styles.slideXReverse)
        setTimeout(() => {
            element.classList.remove(direction(lastView, currentView), styles.slideX)
            element.classList.add(styles.active)
            old.classList.remove(direction(currentView, lastView), styles.slideX)
            old.classList.remove(styles.active)
            old.classList.add(styles.hidden)
        }, duration)
    }

    return styles.wrapper
}