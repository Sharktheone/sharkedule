import {Slot, SlotNames} from "@kanban/column/task/slots/slotTypes"
import RenderTags from "@kanban/column/task/slots/render/renderTag"


type Props = {
    slot: Slot
}

function Render({slot}: Props) {
    switch (slot.type) {
        case SlotNames.TAGS:
            return RenderTags(slot)
        default:
            return null
    }
}