import styles from "./styles.module.scss"

type Props = {
    size?: string
}

export default function Loader({size}: Props) {
    return (
        <div className={styles.wrapper} style={{
            height: size, width: size
        }}>
            <div className={styles.loader} style={{
                height: size, width: size, borderWidth: `calc(${size} / 6)`
            }}>
                Loading...
            </div>
        </div>
    )
}