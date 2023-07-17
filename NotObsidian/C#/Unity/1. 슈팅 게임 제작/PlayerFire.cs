using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerFire : MonoBehaviour
{
    public GameObject bulletFactory;
    public GameObject firePosition;
    // Start is called before the first frame update
    void Start()
    {
        
    }

    // Update is called once per frame
    void Update()
    {
        // 1. 사용자가 발사 버튼을 누르면
        if (Input.GetButtonDown("Fire1"))
        {
            // 2. 총알 공장에서 총알을 만든다
            GameObject bullet = Instantiate(bulletFactory);

            // 3. 총알을 발사한다.
            bullet.transform.position = firePosition.transform.position;
        }
    }
}
