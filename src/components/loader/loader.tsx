import styles from "./styles.module.scss"
import {useColors} from "./colors"

export default function Loader() {
    const {cx, classes} = useColors()
    return (
        <div className={`${styles.loader} ${cx(classes.loader)}`}>
            Loading...
        </div>
    )
}