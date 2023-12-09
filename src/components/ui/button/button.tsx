import styles from "./button.module.scss"
import {HTMLAttributes} from "react"

type props = {
    variant?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark"
    noClickEffect?: boolean
    noHoverEffect?: boolean
    gradient?: boolean
    type?: "button" | "submit" | "reset"
    //color?: "green" | "red" | "blue" | "yellow" | "purple" | "pink" | "orange" | "gray" | "black" | "white" //TODO

} & HTMLAttributes<HTMLButtonElement>

export function Button({variant, gradient, onClick, noClickEffect, noHoverEffect, type, ...props}: props) {

    let vari: string

    switch (variant) {
        case "primary":
            vari = styles.primary
            break
        case "secondary":
            vari = styles.secondary
            break
        case "danger":
            vari = styles.danger
            break
        case "success":
            vari = styles.success
            break
        case "warning":
            vari = styles.warning
            break
        case "info":
            vari = styles.info
            break
        case "light":
            vari = styles.light
            break
        case "dark":
            vari = styles.dark
            break
        default:
            vari = styles.primary
            break

    }

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
        <button onClick={onClickHandler} type={type}
                className={`${styles.button} ${vari} ${gradient ? styles.gradient : ""} ${noHoverEffect ? styles.noHover : ""}`} {...props}/>
    )
}