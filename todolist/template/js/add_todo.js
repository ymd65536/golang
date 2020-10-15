function AddToDo(){
    let ToDoRow = document.getElementsByName("ToDoList");
    let NowTodRows = document.getElementsByName("ToDoRow");
    let RowValue = NowTodRows.length + 1;
  
    let AddRow = `
      <tr name='ToDoRow'>
        <td name='DoCheck'><input type='checkbox' value='${RowValue}'/></td>
        <td name='ToDoText' ><input type='textarea' /></td>
        <td name='ToDoDate'></td>
        <td name='DeleteToDo'><input type='button' value='-' onClick='DeleteToDo(${RowValue})'></td>
      </tr>`
      ToDoRow.item(0).innerHTML = ToDoRow.item(0).innerHTML + AddRow;
  }
  