import styles from "./button.module.scss"

type props = {
    variant?: "primary" | "secondary" | "danger" | "success" | "warning" | "info" | "light" | "dark"
    gradient?: boolean

} & React.HTMLAttributes<HTMLButtonElement>

export default function Button({ variant, gradient, ...props }: props) {
    return (
        <button className={`${styles.button} ${variant} ${gradient ? "gradient" : ""}`} {...props}/>
    )
}