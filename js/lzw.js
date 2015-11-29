var alphabet = ".^:-=+*#%@";

function LZ78 (alphabet){
    var lookup = [{char:0,prev:0}];
  
     
    [].slice.call(alphabet).forEach(function(a){
        var b = a.charCodeAt(0);
        lookup.push({prev:0, char:b});
    }); 

    var rebuild = function(index){
        var word = [];

        var cur = lookup[index];
        while (cur.char != 0) {
            word.push(cur.char);
            cur = lookup[cur.prev];
        }

        return word;
    };

    var decode = function(bytes){
        
        var view = new Uint16Array(bytes);

        var output = [lookup[view[0]].char];

        for(var i = 1; i < view.length; i++){
            var word;

            if (view[i] >= lookup.length){
                var sc = rebuild(view[i-1]);
                lookup.push({char:sc.slice(-1)[0], prev: view[i-1]});
                word = rebuild(view[i]);            
            }else{
                word = rebuild(view[i]);
                lookup.push({char:word[word.length-1],prev:view[i-1]});
            }

            output = output.concat(word.reverse());
        }

        return output;

    };

    return {decode:decode};
}
