import {HTMLAttributes} from "react";

type props = {
    a?: "left" | "center" | "right"
    s?: "small" | "medium" | "large"
    w?: "light" | "regular" | "bold"
    c?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark"
} & HTMLAttributes<HTMLDivElement>


export function Text({a, s, w, c, children, ...props}: props) {
    return (
        <p> {children} </p>
    )
}