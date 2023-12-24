import styles from "./button.module.scss"
import {HTMLAttributes} from "react"

type props = {
    variant?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark" | "green" | "red" | "blue" | "yellow" | "purple" | "pink"  | "indigo" | "cyan" | "teal" | "orange" | "gray" | "black" | "white"
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
        <button onClick={onClickHandler} type={type} data-color={variant ?? "primary"}
                className={`${styles.button} ${gradient ? styles.gradient : ""} ${noHoverEffect ? styles.noHover : ""}`} {...props}/>
    )
}