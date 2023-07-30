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


        // THEORY:
        // show new:
        //  move to overflow (left or right) => hiddenLeft or hiddenRight with translateX
        //  translate to 0
        // remove hiddenLeft or hiddenRight
        // add active
        // hide old:
        // remove active
        // translate to hiddenLeft or hiddenRight
        // add hidden
        // remove hiddenLeft or hiddenRight

        console.log(direction(lastView, currentView))
        element.classList.remove(styles.hidden)
        element.classList.add(direction(currentView, lastView), styles.slideXReverse, styles.active, styles.opacityShow)
        old.classList.add(direction(currentView, lastView), styles.slideXReverse, styles.opacityHide)
        setTimeout(() => {
            element.classList.remove(direction(lastView, currentView), styles.slideXReverse, styles.opacityShow)
            old.classList.remove(direction(currentView, lastView), styles.slideX, styles.opacityHide)
            old.classList.add(styles.hidden)
            old.classList.remove(styles.active)
        }, duration)
    }

    return styles.wrapper
}