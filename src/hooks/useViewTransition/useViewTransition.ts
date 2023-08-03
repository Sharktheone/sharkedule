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

        // I need to add more classes. But maybe it works with just one class per direction
        if (currentIndex < newIndex) {
            // transitioning from right to left
            return styles.hiddenRight
        } else if (currentIndex > newIndex) {
            // transitioning from left to right
            return styles.hiddenLeft
        }
        // the same view, so we don't need to transition
        return ""
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


        //we need to change some classes depending on the direction


        //this is for right to left
        // console.log(direction(currentView, lastView))
        // element.classList.add(styles.hiddenLeft, styles.slideX, styles.active, styles.opacityShow)
        // element.classList.remove(styles.hidden)
        // old.classList.add(direction(lastView, currentView), styles.slideX, styles.opacityHide)
        // setTimeout(() => {
        //     element.classList.remove(direction(lastView, currentView), styles.slideXReverse, styles.opacityShow)
        //     old.classList.remove(direction(currentView, lastView), styles.slideX, styles.opacityHide, styles.slideXReverse)
        //     old.classList.add(styles.hidden)
        //     old.classList.remove(styles.active)
        // }, duration)

        // for left to right i need to figure it out:
        console.log(direction(currentView, lastView))
        element.classList.add(direction(lastView, currentView), styles.slideXReverse, styles.active, styles.opacityShow)
        element.classList.remove(styles.hidden)
        old.classList.add(direction(currentView, lastView), styles.slideX, styles.opacityHide)
        setTimeout(() => {
            element.classList.remove(direction(lastView, currentView), styles.slideXReverse, styles.opacityShow)
            old.classList.remove(direction(currentView, lastView), styles.slideX, styles.opacityHide, styles.slideXReverse)
            old.classList.add(styles.hidden)
            old.classList.remove(styles.active)
        }, duration)
    }

    return styles.wrapper
}