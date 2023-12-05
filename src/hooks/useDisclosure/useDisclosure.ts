import {Simulate} from "react-dom/test-utils"


export default function useDisclosure(initial: boolean = false) {
    return [initial,
        {
            open: () => {
            },
            close: () => {
            },
            toggle: () => {
            },
        },
    ] as const
}