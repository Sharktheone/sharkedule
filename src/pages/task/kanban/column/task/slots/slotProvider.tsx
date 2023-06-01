import {createContext, ReactNode} from "react"


type SlotType = {
    type: "tag" | "priority" | "field" | "assignee" | "subtasks" | "due_date" | "stage" | "progress" | "image"
    // TODO: Add fields for each type

}


type SlotContextType = {
    upperSlot: string
    lowerSlot: string
}

const SlotContext = createContext<SlotContextType | undefined>(undefined)

type Props = {
    children: ReactNode
}

enum SlotTypes {
    UPPER,
    LOWER
}


function SlotProvider({children}: Props) {
    return (
        <SlotContext.Provider value={{upperSlot: "", lowerSlot: ""}}>
            {children}
        </SlotContext.Provider>
    )
}