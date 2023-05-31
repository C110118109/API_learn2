create function add_department_code()
returns text as
$$
declare
    old_id text :=(select  department_code  from departments order by department_code desc limit 1);
    id_number varchar(3) :='001';
    -- datetime char(4) :=substring(cast(now() as varchar),1,4);
    new_id text ;
    num integer;
begin
    if old_id is null then
        new_id:='dept-'||id_number;
        return new_id;
    end if;
    
   
    num :=cast(right(old_id,3) as integer)+1;
    id_number:=
    case
        when num<10 then '00'||num
        when num<10 then '0'||num
		when num<100 then cast(num as text)
        
    end;
    
    
    new_id:='dept-'||id_number;
    return new_id;
end; 
$$
language 'plpgsql';