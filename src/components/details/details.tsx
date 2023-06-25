import {ReactNode} from "react"
import {details} from "@/components/details/types"
import {Drawer, Modal} from "@mantine/core"


type Props = {
    children: ReactNode
    open: boolean
    onClose: () => void
}


//TODO: make changable
let type = details.DRAWER

export default function Details({children, onClose, open}: Props) {
    if (type === details.DRAWER) {
        return (
            <Drawer opened={open} onClose={onClose} position="right"
                    overlayProps={{opacity: 0.15, blur: 1}}
            >
                {children}
            </Drawer>
        )
    } else {
        return (
            <Modal opened={open} onClose={onClose}
                   overlayProps={{opacity: 0.15, blur: 1}}
            >
                {children}
            </Modal>
        )
    }

}