import {ReactNode} from "react"
import {details} from "@/components/details/types"
import {Drawer, Modal} from "@mantine/core"


type Props = {
    children: ReactNode
    open: boolean
    onClose: () => void
    title?: string
    about?: string
}


//TODO: make changable
let type = details.DRAWER

export default function Details({children, onClose, open, about, title}: Props) {
    if (type === details.DRAWER) {
        return (
            <Drawer opened={open} onClose={onClose} position="right"
                    about={about} title={title}
                    overlayProps={{opacity: 0.15, blur: 2}}
            >
                {children}
            </Drawer>
        )
    } else {
        return (
            <Modal opened={open} onClose={onClose}
                   about={about} title={title}
                   overlayProps={{opacity: 0.15, blur: 2}}
            >
                {children}
            </Modal>
        )
    }

}