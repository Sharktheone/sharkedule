type Props = {
    show?: boolean
    text: string
    position?: "top" | "bottom" | "left" | "right"
    element?: HTMLElement
    rect?: DOMRect
}

export function Tooltip({show, text, position, rect, element}: Props) {

    if (!element && !rect) return null
    if (!show) return null


    return (
        <div>
            <h1>Tooltip</h1>
            <p>TODO: Tooltip</p>
        </div>
    )
}