import styles from "./button.module.scss"
import {HTMLAttributes} from "react"
import {ColorModifier} from "@/types/color/color"

type props = {
    variant?: ColorModifier
    noClickEffect?: boolean
    noHoverEffect?: boolean
    gradient?: boolean
    type?: "button" | "submit" | "reset"

} & HTMLAttributes<HTMLButtonElement>

export function Button({variant, gradient, onClick, noClickEffect, noHoverEffect, type, ...props}: props) {

    function onClickHandler(event: any) {
        if (!noClickEffect) {
            event.target.classList.add(styles.clicked)
            setTimeout(() => {
                event.target.classList.remove(styles.clicked)
            }, 100)

        }

        if (onClick) {
            onClick(event)
        }
    }

    return (
        <button onClick={onClickHandler} type={type} data-bg-color={variant ?? "primary"}
                className={`${styles.button} ${gradient ? styles.gradient : ""} ${noHoverEffect ? styles.noHover : ""}`} {...props}/>
    )
}