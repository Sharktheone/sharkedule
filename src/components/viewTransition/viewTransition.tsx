import React, {createRef, ReactElement} from "react"
import {viewRef} from "@/hooks/useViewTransition/useViewTransition"


type props = {
    view: string
    children: ReactElement[]

}

export default function ViewTransition({view, children}: props) {
    console.log(children[0])

    let refs = [] as viewRef[]

    console.log()

    children[0]

    const ref = createRef()

    return (
        children
    )

}