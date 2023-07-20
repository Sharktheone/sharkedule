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
        element.classList.add(styles.hidden, styles[transition ?? ""])
    }

    function show(element: HTMLElement) {
        const old = getOldElement()
        element.classList.remove(styles.hidden)
        if (!old) {
            console.log("no old")
            element.classList.add(styles.active)
            return
        }

        element.classList.add(direction(currentView, lastView), styles.slideX)
        old.classList.add(direction(lastView, currentView), styles.slideX)
        setTimeout(() => {
            element.classList.remove(direction(currentView, lastView), styles.slideX)
            old.classList.remove(direction(lastView, currentView), styles.slideX)
            element.classList.add(styles.active)
            old.classList.remove(styles.active)
            old.classList.add(styles.hidden)
        }, duration)
        //TODO: transition
    }
    return styles.wrapper
}