import React, {useState, ReactElement, useEffect, useRef} from "react"
import useViewTransition, {viewRef} from "@/hooks/useViewTransition/useViewTransition"


type props = {
    view: string
    children: ReactElement[]

}

export default function ViewTransition({view, children}: props) {
    // console.log(children)
    const ref = useRef<HTMLDivElement | null>(null)
    const [lastView, setLastView] = useState<string | null>(null)

    let refs = [] as viewRef[]


    for (let i = 0; i < children.length; i++) {
        refs.push({element: ref.current?.children?.item(i) ?? {} as HTMLElement, id: children[i].props["data-id"]})
    }

    useEffect(() => {
        return () => {
            setLastView(view)
        }
    }, [view])

    let r = useViewTransition(view, lastView, refs)


    return (
        <div ref={(re) => {
            r(re)
            ref.current = re
        }}>
            {children}
        </div>
    )

}