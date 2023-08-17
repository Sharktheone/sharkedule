

type ColorShades = {
    rows: ColorRow[]
}

type ColorRow = {
    color: Color[]
}

type Color = {
    // mainly hsl is used, but we can add rgb if we need
    h: number
    s: number
    l: number
    //
    // r: number
    // g: number
    // b: number
}

export function ColorSelector() {
    // TODO: This is ride now just a note to me...
    // Maybe add some configurable colors that are predefined and also allow user defined colors.
    // For the custom also add places to store them and also allow "single-use" colors, but maybe list them somewhere.
    // Also for the custom colors first when you define them only let them change the hsl h-value, and add a extend button for the whole spectrum

    // (use viewTransition for this one, but maybe let the option, so we can use this as a "popup" variant - user-defined?

    return (
        <div>
            hello
        </div>
    )
}