const token = "token"

function showSuggestion(suggestion) {
  console.log(suggestion);
}

$(".delivery").suggestions({
    token: token,
    type: "ADDRESS",
    onSelect: showSuggestion
});

$(".fullname").suggestions({
  token: token,
  type: "NAME",
  onSelect: showSuggestion
});


$(".email").suggestions({
    token: token,
    type: "EMAIL",
    onSelect: showSuggestion
});

function join(arr /*, separator */) {
  var separator = arguments.length > 1 ? arguments[1] : ", ";
  return arr.filter(function(n){return n}).join(separator);
}

function showClientSuggestion(suggestion) {
    console.log(suggestion);
    var data = suggestion.data;
     if (!data)
       return;
      
     if (data.name) {
      $(".c-name-full").val(data.name.full_with_opf || "");
    }
       
     $(".c-inn-kpp").val(join([data.inn, data.kpp], " / "));
    
     if (data.address) {
      var address = "";
      if (data.address.data.qc == "0") {
         address = join([data.address.data.postal_code, data.address.value]);
      } else {
         address = data.address.data.source;
       }
      $(".c-company-address").val(address);
    }
}

// наименование компании
$(".c-party").suggestions({
  token: token,
  type: "PARTY",
  count: 5,
  onSelect: showClientSuggestion
});

function showProviderSuggestion(suggestion) {
  console.log(suggestion);
  var data = suggestion.data;
   if (!data)
     return;
    
   if (data.name) {
    $(".p-name-full").val(data.name.full_with_opf || "");
  }
     
   $(".p-inn-kpp").val(join([data.inn, data.kpp], " / "));
  
   if (data.address) {
    var address = "";
    if (data.address.data.qc == "0") {
       address = join([data.address.data.postal_code, data.address.value]);
    } else {
       address = data.address.data.source;
     }
    $(".p-company-address").val(address);
  }
}

$(".p-party").suggestions({
  token: token,
  type: "PARTY",
  count: 5,
  onSelect: showProviderSuggestion
});
