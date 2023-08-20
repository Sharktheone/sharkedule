type ColorShades = {
    colors: Color[]
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

    function getColors(): ColorShades[] {
        const num = 8
        const startHue = 25
        const s = 75
        const l = 50

        let shades = [] as ColorShades[]

        for (let h = startHue; h < 360 + startHue; h += (360 / num)) {
            let color: Color = {
                h: h,
                s: s,
                l: l,
            }
            shades.push({
                colors: [color]
            })
        }
        return shades
    }

    function get(color: Color) {
        return `hsl(${color.h}deg, ${color.s}%, ${color.l}%)`
    }


    return (
        <div>
            HELLOW
            {getColors().map(shade => (
                    <div>
                        {
                            shade.colors.map(color => (
                                <div style={{
                                    width: "1rem",
                                    height: ".5rem",
                                    backgroundColor: get(color)
                                }}/>
                            ))
                        }
                    </div>
                ))
            }
        </div>
    )
}