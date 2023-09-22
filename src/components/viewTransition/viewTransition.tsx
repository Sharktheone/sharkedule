import React, {createRef, ReactElement, useEffect, useRef} from "react"
import useViewTransition, {viewRef} from "@/hooks/useViewTransition/useViewTransition"
import {usePrevious} from "@mantine/hooks"


type props = {
    view: string
    children: ReactElement[]

}

export default function ViewTransition({view, children}: props) {
    const ref = useRef<HTMLDivElement | null>(null)

    let refs = [] as viewRef[]


    for (let i = 0; i < children.length; i++) {
        refs.push({element: ref.current?.children?.item(i) ?? new Element, id: children[i].props["data-id"]})
    }

    const lastView = usePrevious(view) ?? view

    let r = useViewTransition(view, lastView, refs)

    r(ref.current)


    return (
        <div ref={ref}>
            {children}
        </div>
    )

}