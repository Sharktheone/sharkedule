type Props = {
    format?: "hex" | "rgb" | "hsl"
    onChange?: (value: string) => void
}


export function ColorPicker({format}: Props) {

    format ||= "hex"

    return (
        <div/>
    )
}