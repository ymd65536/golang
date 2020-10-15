// やることを追加する
function InitPage(){
    let RowValue = 1;
    let ToDoRow = document.getElementsByName("ToDoList");
    let TableHead =`
      <th>チェック</th>
      <th>やりたいこと</th>
      <th>終わった日</th>
      <th>完了</th>
    `
    let AddRow = `
      <tr name='ToDoRow'>
        <td name='DoCheck'><input type='checkbox' value='1'/></td>
        <td name='ToDoText' ><input type='textarea' /></td>
        <td name='ToDoDate'></td>
        <td name='DeleteToDo'><input type='button' value='-' onClick='DeleteToDo(${RowValue})'></td>
      </tr>`
      ToDoRow.item(0).innerHTML = TableHead + AddRow;
  }
  