import React, {useState, ReactElement, useEffect, useRef} from "react"
import useViewTransition, {viewRef} from "@/hooks/useViewTransition/useViewTransition"
import {usePrevious} from "@mantine/hooks"


type props = {
    view: string
    children: ReactElement[]

}

export default function ViewTransition({view, children}: props) {
    // console.log(children)
    const ref = useRef<HTMLDivElement | null>(null)

    let refs = [] as viewRef[]


    for (let i = 0; i < children.length; i++) {
        refs.push({element: ref.current?.children?.item(i) ?? {} as HTMLElement, id: children[i].props["data-id"]})
    }

    let lastView = usePrevious(view)

    let r = useViewTransition(view, lastView, refs)


    return (
        <div ref={(re) => {
            ref.current = re
            r(re)
        }}>
            {children}
        </div>
    )

}