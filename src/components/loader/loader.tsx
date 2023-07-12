import styles from "./styles.module.scss"
import {useColors} from "./colors"

type Props = {
    size?: string
}

export default function Loader({size}: Props) {
    const {cx, classes} = useColors()
    return (
        <div className={`${styles.loader} ${cx(classes.loader)}`} style={{
            height: size, width: size, borderWidth: `calc(${size} / 10)`
        }}>
            Loading...
        </div>
    )
}