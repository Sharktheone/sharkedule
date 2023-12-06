import {useState} from "react"


export function useDisclosure(initial: boolean = false) { //Is this hook really useful?

    const [isOpen, setIsOpen] = useState(initial)

    return [initial,
        {
            open: () => {
                setIsOpen(true)
            },
            close: () => {
                setIsOpen(false)
            },
            toggle: () => {
                setIsOpen(!isOpen)
            },
        },
    ] as const
}