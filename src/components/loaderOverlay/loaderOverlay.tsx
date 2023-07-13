import styles from './styles.module.scss'
import Loader from "@/components/loader/loader"
import {ReactNode} from "react"

type Props = {
    loading: boolean
    children: ReactNode
}

export default function LoaderOverlay({loading, children}: Props) {
    return (
        <div className={styles.overlay}>
            {children}
            {loading &&
                <div className={styles.loader}>
                    <Loader size="30px"/>
                </div>
            }
        </div>
    )


}