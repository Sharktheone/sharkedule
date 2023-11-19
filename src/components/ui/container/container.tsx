
type props = {
    noBorder?: boolean
    noBackground?: boolean
    noPadding?: boolean
} & React.HTMLAttributes<HTMLDivElement>


export function Container({noBorder, noBackground, noPadding, ...props}: props) {
    return (
        <div
            className={`container ${noBorder ? "no-border" : ""} ${noBackground ? "no-background" : ""} ${noPadding ? "no-padding" : ""}`}
            {...props}/>
    )
}