import {useContext} from "react"
import {SlotContext} from "@kanban/column/task/slots/slotContext"
import styles from "@kanban/column/task/slots/styles.module.scss"
import {Slot} from "@kanban/column/task/slots/slotTypes"
import Render from "@kanban/column/task/slots/render/render"


export default function LowerSlot() {
    const context = useContext(SlotContext)
    let lowerSlot = context?.lowerSlot
    if (!lowerSlot) return null

    let allUndefined = true
    lowerSlot.forEach((slot: Slot) => {
        if (slot.value !== undefined  && allUndefined) allUndefined = false
    })

    if (allUndefined) return null

    if (lowerSlot.length === 0) return null


    return (
        <div className={styles.slot}>
            {lowerSlot.map((slot: Slot) => (
                <Render key={slot.type} slot={slot}/>
            ))}
        </div>
    )
}