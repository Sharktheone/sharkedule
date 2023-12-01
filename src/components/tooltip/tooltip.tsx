
type Props = {
    show?: boolean
    text: string
    position?: "top" | "bottom" | "left" | "right"
}

export default function Tooltip({show, text, position}: Props) {
    return (
        <div>
            <h1>Tooltip</h1>
            <p>TODO: Tooltip</p>
        </div>
    )
}