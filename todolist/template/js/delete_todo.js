//やることを削除する
function DeleteToDo(RowValue){
    let SelectRowValue  = 0;
    let DoChecks = document.getElementsByName('DoCheck');
    let DoCheck = null;
    let DoCheckCount = 0;
  
    for(DoCheckCount = 0;DoCheckCount <DoChecks.length;DoCheckCount++){
      DoCheck = DoChecks.item(DoCheckCount)
      SelectRowValue = DoCheck.firstElementChild.getAttribute("value")
      if (SelectRowValue == RowValue ){
        console.log(SelectRowValue)
        DoCheck.parentElement.remove();
        break;
      }
    }
  }