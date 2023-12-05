import {useEffect, useState} from "react"
import styles from "./styles.module.scss"


export type viewRef = {
    id: string
    element: HTMLElement | Element
}


export function useViewTransition(currentView: string, lastView: string, viewList: viewRef[], duration: number = 200, transition?: string, timingFunction?: string) {
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
    const [ref, setRef] = useState<HTMLDivElement | null>(null)

    useEffect(() => {
        ref?.classList.add(styles.wrapper)
        ref?.style.setProperty("--duration", duration.toString())
    }, [ref])


    function direction(currentView: string, newView: string) {
        const currentIndex = viewList.findIndex(({id}) => id === currentView)
        const newIndex = viewList.findIndex(({id}) => id === newView)

        // I need to add more classes. But maybe it works with just one class per direction
        if (currentIndex < newIndex) {
            // transitioning from right to left
            return styles.right
        } else if (currentIndex > newIndex) {
            // transitioning from left to right
            return styles.left
        }
        // the same view, so we don't need to transition
        return styles.none
    }

    function getOldElement() {
        if (!lastView) return
        return viewList.find(({id}) => id === lastView)?.element
    }

    function hide(element: HTMLElement | Element) {
        element?.classList?.remove(styles.active)
        element?.classList?.add(styles.hidden)
    }

    function show(element: HTMLElement | Element) {
        const old = getOldElement()
        element?.classList?.remove(styles.hidden)
        if (!old) {
            element?.classList?.add(styles.active)
            return
        } else if (old === element) {
            element?.classList?.add(styles.active)
            return
        }

        element?.classList?.add(direction(currentView, lastView), styles.active)
        element?.classList?.remove(styles.hidden)
        old?.classList?.add(direction(currentView, lastView))
        setTimeout(() => {
            element?.classList?.remove(direction(lastView, currentView), styles.left, styles.right)
            old?.classList?.remove(direction(currentView, lastView), styles.active, styles.left, styles.right)
            old?.classList?.add(styles.hidden)
        }, duration)
    }

    return setRef
}