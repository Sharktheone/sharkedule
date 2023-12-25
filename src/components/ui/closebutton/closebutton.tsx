import {Button} from "@/components/ui"
import {HTMLAttributes} from "react"
import {IconX} from "@tabler/icons-react"
import styles from "./styles.module.scss"


type Props = {} & HTMLAttributes<HTMLButtonElement>

export function CloseButton({...props}: Props) {
    return (
        <div>
            <Button {...props} className={styles.close}>
                <IconX/>
            </Button>
        </div>
    )
}