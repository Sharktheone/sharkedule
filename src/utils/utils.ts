export default function calculateBrightnessHEX(hex: string) {
    let R, G, B: number
    if (hex.length === 7) {
        R = parseInt(hex.slice(1, 3), 16)
        G = parseInt(hex.slice(3, 5), 16)
        B = parseInt(hex.slice(5, 7), 16)
    } else if (hex.length === 4) {
        R = parseInt(hex.slice(1, 2), 16)
        G = parseInt(hex.slice(2, 3), 16)
        B = parseInt(hex.slice(3, 4), 16)
    } else {
        throw new Error("Invalid HEX")
    }

    return (R * 299 + G * 587 + B * 114) / 1000

}