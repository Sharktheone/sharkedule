import {createStyles, Group, rem} from '@mantine/core';
import test_data from "./test_data.json"
import {DragDropContext, DropResult, ResponderProvided} from "react-beautiful-dnd"
import Column from "./column/column";
import {useState} from "react"
import {kanbanBoardType} from "./types"
export default function Kanban() {
    const [board, setBoard] = useState<kanbanBoardType>(test_data as unknown as kanbanBoardType)
    // const items = state.map((item, index) => (
    //     <Draggable key={item.symbol} index={index} draggableId={item.symbol}>
    //         {(provided, snapshot) => (
    //             <div
    //                 className={cx(classes.item, {[classes.itemDragging]: snapshot.isDragging})}
    //                 {...provided.draggableProps}
    //                 {...provided.dragHandleProps}
    //                 ref={provided.innerRef}
    //             >
    //                 <Text className={classes.symbol}>{item.symbol}</Text>
    //                 <div>
    //                     <Text>{item.name}</Text>
    //                     <Text color="dimmed" size="sm">
    //                         Position: {item.position} • Mass: {item.mass}
    //                     </Text>
    //                 </div>
    //             </div>
    //         )}
    //     </Draggable>
    // ));


    return (
        <Group position="center" noWrap={true}>
            {board.columns?.map((column) => (
                <Column key={column.uuid} column={column}/>
            ))}
        </Group>
)
}

