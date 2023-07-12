import styles from "./styles.module.scss"
import {useColors} from "./colors"

type Props = {
    size?: string
}

export default function Loader({size}: Props) {
    const {cx, classes} = useColors()
    return (
        <div className={styles.wrapper} style={{
            height: size, width: size
        }}>
            <div className={`${styles.loader} ${cx(classes.loader)}`} style={{
                height: size, width: size, borderWidth: `calc(${size} / 6)`
            }}>
                Loading...
            </div>
        </div>
    )
}